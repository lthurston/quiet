package util

import (
	"os/user"
	"regexp"
)


/// / getHomeDir returns the user's home directory -- this belongs elsewhere
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	return usr.HomeDir
}

// Detilde replaces a tilde with the full path to the user's home directory
func Detilde(filename string) string {
	return string(regexp.MustCompile("^~").ReplaceAll([]byte(filename), []byte(GetHomeDir())))
}