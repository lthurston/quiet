package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/lthurston/quiet/util"

	goTomlConfig "github.com/stvp/go-toml-config"
)

var (
	configFile                  = goTomlConfig.String("file", util.GetHomeDir()+"/.ssh/config")
	configListFields            = goTomlConfig.String("list.fields", "User, Hostname")
	configNewFrom               = goTomlConfig.String("new.from", "")
	configExportFilenameOptions = goTomlConfig.String("export.filenameOptions", "IdentityFile")
	configHostTemplate          = goTomlConfig.String("template", `
Host {{.Name}}{{if .Aliases}}{{range .Aliases}} {{.}}{{end}}{{end}}
{{range .Options }}	{{.}}
{{end}}`)
)

// quietConfig stores the location of the quiet config, defaults to "~/.quiet"
var quietConfig = ""

// Parseable tells us if the config file can be read / parsed
func Parseable(quietConfig string) bool {
	return goTomlConfig.Parse(quietConfig) == nil
}

// GetConfigFile returns `file`
func GetConfigFile() string {
	return *configFile
}

// GetConfigListFields returns `list.show`
func GetConfigListFields() string {
	return *configListFields
}

// GetConfigNewFrom returns `new.from`
func GetConfigNewFrom() string {
	return *configNewFrom
}

// GetConfigExportFileOptions
func GetConfigExportFilenameOptions() string {
	return *configExportFilenameOptions
}

// GetConfigTemplate returns the template
func GetConfigTemplate() string {
	return *configHostTemplate
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
		quietConfig = util.GetHomeDir() + "/.quiet"
	}

	if _, err := os.Stat(quietConfig); err != nil {
		if os.IsNotExist(err) {
			create(quietConfig)
			fmt.Println("Config file created at " + quietConfig + " with __ALL__ my favorite defaults!")
			return
		}
	}

	if !Parseable(quietConfig) {
		fmt.Println("Your .quiet configuration exists, but it's not parsable. :(")
		os.Exit(1)
	}
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
		`file = "` + *configFile + `"

[list]
fields = "` + *configListFields + `"

[new]
from = "` + *configNewFrom + `"

[export]
filenameOptions = "`+  *configExportFilenameOptions + `"`,
	)
	return
}
