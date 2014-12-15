package config

import (
	"fmt"
	"io/ioutil"
	"os/user"

	goTomlConfig "github.com/stvp/go-toml-config"
)

var (
	configFile       = goTomlConfig.String("file", getHomeDir()+"/.ssh/config")
	configListFields = goTomlConfig.String("list.fields", "User, Hostname")
	configListFormat = goTomlConfig.String("list.format", "%-20s %s")
)

// quietConfig stores the location of the quiet config, defaults to "~/.quiet"
var quietConfig = ""

// Parseable tells us if the config file can be read / parsed
func Parseable(quietConfig string) bool {
	return goTomlConfig.Parse(quietConfig) != nil
}

// GetConfigFile returns `file`
func GetConfigFile() string {
	return *configFile
}

// GetConfigListFields returns `list.show`
func GetConfigListFields() string {
	return *configListFields
}

// GetConfigListFormat returns `list.format`
func GetConfigListFormat() string {
	return *configListFormat
}

// GetConfigMap returns a map of all configuration values
func GetConfigMap() map[string]string {
	configMap := make(map[string]string)
	configMap["file"] = *configFile
	return configMap
}

// SetQuietConfig sets the location of the quiet config file
func SetQuietConfig(qc string) {
	quietConfig = qc
	parse()
}

func init() {
	parse()
}

// parse reads the config, or creates one with defaults; an alternate file location
// for .quiet can be passed in, but this is likely not all that useful
func parse() {
	if quietConfig == "" {
		quietConfig = getHomeDir() + "/.quiet"
	}

	if !Parseable(quietConfig) {
		create(quietConfig)
		fmt.Println("Config file created at " + quietConfig + " with __ALL__ my favorite defaults!")
	}
}

// getHomeDir returns the user's home directory -- this belongs elsewhere
func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	return usr.HomeDir
}

// create will write a configuration file
func create(file string) {
	configBytes := getQuietConfigBytes()

	err := ioutil.WriteFile(file, configBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// getQuietConfigBytes returns a bytestring of the config values -- probably
// should use the GetConfigMap
func getQuietConfigBytes() (bytes []byte) {
	bytes = []byte(
		`file = ` + *configFile,
	)
	return
}
