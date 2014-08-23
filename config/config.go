package config

import (
	"github.com/stvp/go-toml-config"
	"os/user"
	"fmt"
)

var (
	File            = config.String("file", "~/.ssh/config")
)

func Parse() {
	usr, err := user.Current()
    if err != nil {
        panic(err)
    }

	if err := config.Parse(usr.HomeDir + "/.quiet"); err != nil {
		create()
	}
}

func create() {
	fmt.Println("Let's make that config file")
}