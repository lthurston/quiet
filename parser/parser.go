package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type host struct {
	Name      string
	Aliases   []string
	Config    map[string]string
	StartLine int
	EndLine   int
}

func makeHost() host {
	i := host{}
	i.Config = make(map[string]string)
	return i
}

func (host *host) addConfigFromString(line string) {
	line = strings.TrimSpace(line)
	sepIndex := strings.IndexAny(line, " \t")
	config, value := line[0:sepIndex], line[sepIndex+1:]
	host.Config[config] = value
}

// HostsCollection is a collection of hosts from the config file!
type HostsCollection struct {
	Hosts []host
}

var ignoreLineRegexes = []string{
	`^\s*$`,
	`^\s*#`,
	`^\s*Host \*`,
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

// fromScanner populates the host from a bufio.Scanner
func (hosts *HostsCollection) fromScanner(s *bufio.Scanner) {
	skippableLine := getLineMatcher(ignoreLineRegexes)
	hostLine := getLineMatcher([]string{`^Host.+$`})
	//configLine := getLineMatcher([]string{`^\s+.+\s.+$`})

	// tmp := []string{}

	foundFirstHostLine := false

	// Remove garbage, only start tracking when we hit a "Host" line
	lastNonSkippableLine := 0
	lineIndex := 0

	host := makeHost()
	for s.Scan() {
		lineIndex++
		line := s.Text()
		if !skippableLine(line) {
			if hostLine(line) {
				if foundFirstHostLine {
					host.EndLine = lastNonSkippableLine
					hosts.Hosts = append(hosts.Hosts, host)
				}
				foundFirstHostLine = true
				host = makeHost()
				host.StartLine = lineIndex
				hostNames := getHostNames(line)
				host.Name = hostNames[0]
				host.Aliases = hostNames[1:]
			} else {
				if foundFirstHostLine {
					host.addConfigFromString(line)
				}
			}
			lastNonSkippableLine = lineIndex
		}
	}

	if foundFirstHostLine {
		host.EndLine = lastNonSkippableLine
		hosts.Hosts = append(hosts.Hosts, host)
	}
}

// FindHostByName finds a host by name (not by host aliases!)
func (hosts HostsCollection) FindHostByName(name string) (host, bool) {
	for _, host := range hosts.Hosts {
		if host.Name == name {
			return host, true
		}
	}
	return makeHost(), false
}
