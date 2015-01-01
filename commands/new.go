package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/parser"

	"github.com/spf13/cobra"
)

type inputValidator func(string) error

var newFrom, newName string
var newSkipInteractive bool

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Appends new host using other host as template",
	Long:  `New appends a new host to your SSH configuration based on other host`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("from: ", newFrom)
		fmt.Println("new: ", newName)
		fmt.Println("skip-interactive: ", newSkipInteractive)

		from := config.GetConfigNewFrom()
		if len(newFrom) > 0 {
			from = newFrom
		}

		hosts := parser.HostsCollection{}
		hosts.ReadFromFile(config.GetConfigFile())

		if len(newName) == 0 {
			validator := makeNewHostnameValidator(hosts)
			newName = getNewHostname(newFrom, validator)
		}

		if host, found := hosts.FindHostByName(from); found {
			host.Name = newName
			newHostSnippet := host.RenderSnippet()
			fmt.Println(newHostSnippet)
		} else {
			fmt.Println("Couldn't find host or empty: \"" + from + "\"")
		}
	},
}

func makeNewHostnameValidator(hosts parser.HostsCollection) inputValidator {
	return func(value string) error {
		var err error
		if match, _ := regexp.MatchString("^[\\.\\-a-zA-Z0-9]+$", value); !match {
			return errors.New("Host name contains illegal character(s)")
		}

		if _, found := hosts.FindHostByName(value); found {
			return errors.New("A host by that name already exists.")
		}
		return err
	}
}

func getNewHostname(newFrom string, validator inputValidator) string {
	i := 0
	var err error
	for err != nil || i < 1 {
		if i > 0 {
			fmt.Println("Error: ", err, " Try again.")
		}
		newName = inputWithDefault("Name: ", newFrom)
		err = validator(newName)
		i++
	}

	return newName
}

func inputWithDefault(prompt, value string) string {
	fmt.Print(prompt)
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')

	for err != nil {
		fmt.Println("Got error: ", err, " Try again.")
		fmt.Print(prompt)
		input, err = r.ReadString('\n')
	}

	input = input[:len(input)-1]

	if input == "" {
		input = value
	}

	return input
}
