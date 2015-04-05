package host

import (
	"bufio"
	"regexp"
	"strings"
)

var ignoreLineRegexes = []string{
	`^\s*$`,
	`^\s*#`,
	`^\s*Host \*\s*$`,
}

// fromScanner populates the host from a bufio.Scanner
func fromScanner(hosts *HostsCollection, s *bufio.Scanner) {
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

// A function that you pass a string to, and it will return true if
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
