package config

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/stvp/go-toml-config"
)

var sshFile = config.String("file", "~/.ssh/config")
var strategy = config.String("strategy", "append")

// Parse reads the config, or creates one with defaults
func Parse(quietConfig string) {
	if quietConfig == "" {
		quietConfig = getHomeDir() + "/.quiet"
	}

	if Parseable(quietConfig) {
		create(quietConfig)
		fmt.Println("Config file created at " + quietConfig + " with __ALL__ my favorite defaults!")
	}
}

// Parseable tells us if the config file can be read / parsed
func Parseable(quietConfig string) bool {
	return config.Parse(quietConfig) != nil
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	return usr.HomeDir
}

func create(file string) {
	configBytes := []byte(
		`file = ` + *sshFile +
			`strategy = ` + *strategy,
	)

	err := ioutil.WriteFile(file, configBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// GetSSHFile returns a config value
func GetSSHFile() string {
	return *sshFile
}

// GetStrategy returns a config value
func GetStrategy() string {
	return *strategy
}
