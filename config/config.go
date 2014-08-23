package config

import (
	"github.com/stvp/go-toml-config"
	"os/user"
	"fmt"
	"io/ioutil"
)

var (
	File            = config.String("file", "~/.ssh/config")
)

func Parse() {
    quietConfig := getHomeDir() + "/.quiet"
	if err := config.Parse(quietConfig); err != nil {
		create(quietConfig)
		fmt.Println("Config file created at " + quietConfig + " with __ALL__ my favorite defaults!")
	}
}

func getHomeDir() string {
	usr, err := user.Current()
    if err != nil {
        panic(err)
    }

	return usr.HomeDir 
}

func create(file string) {
	configBytes := []byte(`file = ` + *File)
	err := ioutil.WriteFile(file, configBytes, 0644)
	if err != nil {
		panic(err)
	}
}