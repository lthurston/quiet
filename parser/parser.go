package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	h "github.com/lthurston/quiet/host"
)

// HostsCollection is a collection of hosts from the config file!
type HostsCollection struct {
	Hosts []h.Host
}

// Count counts
func (hosts HostsCollection) Count() int {
	return len(hosts.Hosts)
}

// ReadFromFile loads a HostsCollection with stuff from a fil
func (hosts *HostsCollection) ReadFromFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	hosts.fromScanner(scanner)
}

// ReadFromString is a probably unnecessary wrapper around strings.NewReader()
func (hosts *HostsCollection) ReadFromString(contents string) {
	r := strings.NewReader(contents)

	scanner := bufio.NewScanner(r)
	hosts.fromScanner(scanner)
}

// fromScanner populates the host from a bufio.Scanner
func (hosts *HostsCollection) fromScanner(s *bufio.Scanner) {
	skippableLine := getLineMatcher(ignoreLineRegexes)
	hostLine := getLineMatcher([]string{`^Host.+$`})

	foundFirstHostLine := false

	// Remove garbage, only start tracking when we hit a "Host" line
	lastNonSkippableLine := 0
	lineIndex := 0

	host := h.MakeHost()
	for s.Scan() {
		lineIndex++
		line := s.Text()
		if !skippableLine(line) {
			if hostLine(line) {
				if foundFirstHostLine {
					host.SetEndLine(lastNonSkippableLine)
					hosts.Hosts = append(hosts.Hosts, host)
				}
				foundFirstHostLine = true
				host = h.MakeHost()
				host.SetStartLine(lineIndex)
				hostNames := getHostNames(line)
				host.SetName(hostNames[0])
				host.SetAliases(hostNames[1:])
			} else {
				if foundFirstHostLine {
					host.AddConfigFromString(line)
				}
			}
			lastNonSkippableLine = lineIndex
		}
	}

	if foundFirstHostLine {
		host.SetEndLine(lastNonSkippableLine)
		hosts.Hosts = append(hosts.Hosts, host)
	}
}

// FindHostByName finds a host by name (not by host aliases!)
func (hosts HostsCollection) FindHostByName(name string) (h.Host, bool) {
	for _, host := range hosts.Hosts {
		if host.GetName() == name {
			return host, true
		}
	}
	return h.MakeHost(), false
}

var ignoreLineRegexes = []string{
	`^\s*$`,
	`^\s*#`,
	`^\s*Host \*`,
}

// An function that you pass a string to, and it will return true if
// the string doesn't match one of the ignoreLineRegexes
type lineMatcher func(input string) bool

// ignoreLines returns an ignoreLineFunc
func getLineMatcher(regexStrings []string) lineMatcher {
	regexes := []*regexp.Regexp{}
	for _, regexString := range regexStrings {
		regexes = append(regexes, regexp.MustCompile(regexString))
	}
	return func(input string) bool {
		for _, regex := range regexes {
			if regex.Match([]byte(input)) {
				return true
			}
		}
		return false
	}
}

func getHostNames(line string) []string {
	return strings.Fields(line)[1:]
}
