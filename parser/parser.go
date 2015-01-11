package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	h "github.com/lthurston/quiet/host"
)

// HostPosition keeps track of a set of host configuration and the start and
// end lines within the HostsCollection
type HostPosition struct {
	h.Host
	startLine int
	endLine   int
}

// MakeHostPosition makes a new one of these
func MakeHostPosition() HostPosition {
	return HostPosition{Host: h.MakeHost()}
}

// SetStartLine sets start line
func (hostPosition *HostPosition) SetStartLine(startLine int) {
	hostPosition.startLine = startLine
}

// StartLine get startLine
func (hostPosition HostPosition) StartLine() int {
	return hostPosition.startLine
}

// SetEndLine sets the end line
func (hostPosition *HostPosition) SetEndLine(endLine int) {
	hostPosition.endLine = endLine
}

// EndLine gets the end line
func (hostPosition HostPosition) EndLine() int {
	return hostPosition.endLine
}

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

	host := MakeHostPosition()
	for s.Scan() {
		lineIndex++
		line := s.Text()
		if !skippableLine(line) {
			if hostLine(line) {
				if foundFirstHostLine {
					host.SetEndLine(lastNonSkippableLine)
					hosts.HostPositions = append(hosts.HostPositions, host)
				}
				foundFirstHostLine = true
				host = MakeHostPosition()
				host.SetStartLine(lineIndex)
				hostNames := getHostNames(line)
				host.SetName(hostNames[0])
				host.SetAliases(hostNames[1:])
			} else {
				if foundFirstHostLine {
					host.AddOptionFromString(line)
				}
			}
			lastNonSkippableLine = lineIndex
		}
	}

	if foundFirstHostLine {
		host.SetEndLine(lastNonSkippableLine)
		hosts.HostPositions = append(hosts.HostPositions, host)
	}
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
	if fields := strings.Fields(line)[1:]; len(fields) > 0 {
		return fields
	}
	return []string{""}
}
