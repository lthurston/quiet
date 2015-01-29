package host


import (
	"os"
	"bufio"
	"strings"
)

// HostsCollection is a collection of hosts from the config file!
type HostsCollection struct {
	HostPositions []HostPosition
}

// GetIndex gets the host position at an index
func (hosts HostsCollection) GetIndex(i int) HostPosition {
	return hosts.HostPositions[i]
}

// Count counts
func (hosts HostsCollection) Count() int {
	return len(hosts.HostPositions)
}

// ReadFromFile loads a HostsCollection with stuff from a fil
func (hosts *HostsCollection) ReadFromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fromScanner(hosts, scanner)
}

// ReadFromString is a probably unnecessary wrapper around strings.NewReader()
func (hosts *HostsCollection) ReadFromString(contents string) {
	r := strings.NewReader(contents)

	scanner := bufio.NewScanner(r)
	fromScanner(hosts, scanner)
}

// FindHostByName finds a host by name (not by host aliases!)
func (hosts HostsCollection) FindHostByName(name string) (HostPosition, bool) {
	for _, host := range hosts.HostPositions {
		if host.Name() == name {
			return host, true
		}
	}
	return MakeHostPosition(), false
}


