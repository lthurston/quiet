package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/host"
	"github.com/lthurston/quiet/writer"

	"github.com/spf13/cobra"
)

type inputValidator func(string) error

var newFrom, newName string
var newSkipInteractive, newStdout bool

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Appends new host using other host as template",
	Long:  `New appends a new host to your SSH configuration based on other host`,
}

func init() {
	newCmd.Flags().StringVarP(&newFrom, "from", "f", "", "host to use as template")
	newCmd.Flags().StringVarP(&newName, "name", "n", "", "new host name")
	newCmd.Flags().BoolVarP(&newSkipInteractive, "skip-interactive", "s", false, "just copy; don't allow interactive")
	newCmd.Flags().BoolVarP(&newStdout, "stdout", "o", false, "output to stdout rather than appending SSH config file")
	newCmd.Run = new
}

func new(cmd *cobra.Command, args []string) {
	from := config.GetConfigNewFrom()
	if len(newFrom) > 0 {
		from = newFrom
	}

	hosts := host.HostsCollection{}
	hosts.ReadFromFile(config.GetConfigFile())
	if host, found := hosts.FindHostByName(from); found {
		fmt.Println("Copying host \"" + from + "\"")
		host.SetName(newName)
		host.SetAliases([]string{})

		if !newSkipInteractive {
			if len(newName) == 0 {
				newName = getNewHostname(newFrom, makeNewHostnameValidator(hosts))
			}
			host.SetOptions(getNewOptionValues(host.Options(), makeOptionValueValidator()))
		}

		host.SetName(newName)
		newHostSnippet := host.String()
		if newStdout {
			fmt.Println(newHostSnippet)
		} else {
			writer.Append(newHostSnippet)
		}

	} else {
		fmt.Println("Couldn't find host or empty: \"" + from + "\"")
	}
}

func hostnameRegexValidator(value string) bool {
	match, _ := regexp.MatchString("^[\\.\\-a-zA-Z0-9]+$", value)
	return match
}

func configRegexValidator(value string) bool {
	match, _ := regexp.MatchString("^[\\.\\-a-zA-Z0-9 ~\\/_]+$", value)
	return match
}

func makeOptionValueValidator() inputValidator {
	return func(value string) error {
		var err error

		if !configRegexValidator(value) {
			return errors.New("Host name contains illegal character(s)")
		}

		return err
	}
}

func makeNewHostnameValidator(hosts host.HostsCollection) inputValidator {
	return func(value string) error {
		var err error
		if !hostnameRegexValidator(value) {
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

func getNewOptionValues(options []host.Option, validator inputValidator) []host.Option {
	var newOption host.Option
	newOptions := []host.Option{}
	for _, option := range options {
		newOption = host.MakeOption(
			option.Keyword(),
			inputWithDefault(option.Keyword()+" [default is \""+option.Argument()+"\"]: ", option.Argument()),
		)
		newOptions = append(newOptions, newOption)
	}
	return newOptions
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
