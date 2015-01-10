package parser

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
	textTemplate "text/template"
)

type host struct {
	name      string
	aliases   []string
	config    map[string]string
	startLine int
	endLine   int
}

func (host *host) addConfigFromString(line string) {
	line = strings.TrimSpace(line)
	sepIndex := strings.IndexAny(line, " \t")
	config, value := line[0:sepIndex], line[sepIndex+1:]
	host.config[config] = value
}

func (host *host) SetName(name string) {
	host.name = name
}

func (host host) GetName() string {
	return host.name
}

func (host *host) SetAliases(aliases []string) {
	host.aliases = aliases
}

func (host host) GetAliases() []string {
	return host.aliases
}

func (host *host) SetConfig(config map[string]string) {
	host.config = config
}

func (host host) GetConfig() map[string]string {
	return host.config
}

func (host *host) SetStartLine(startLine int) {
	host.startLine = startLine
}

func (host host) GetStartLine() int {
	return host.startLine
}

func (host *host) SetEndLine(endLine int) {
	host.endLine = endLine
}

func (host host) GetEndLine() int {
	return host.endLine
}

// RenderSnippet renders a host snippet
func (host host) RenderSnippet() string {
	tmpl, err := textTemplate.New("snip").Parse(`
Host {{.GetName}}{{if .GetAliases}}{{range .GetAliases}} {{.}}{{end}}{{end}}
	{{range $key, $value := .GetConfig	 }}{{$key}} {{$value}}
{{end}}
`)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, host)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// MakeHost returns a host!
func MakeHost() host {
	i := host{}
	i.SetConfig(make(map[string]string))
	return i
}

// HostsCollection is a collection of hosts from the config file!
type HostsCollection struct {
	Hosts []host
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

	host := MakeHost()
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
				host = MakeHost()
				host.SetStartLine(lineIndex)
				hostNames := getHostNames(line)
				host.SetName(hostNames[0])
				host.SetAliases(hostNames[1:])
			} else {
				if foundFirstHostLine {
					host.addConfigFromString(line)
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
func (hosts HostsCollection) FindHostByName(name string) (host, bool) {
	for _, host := range hosts.Hosts {
		if host.GetName() == name {
			return host, true
		}
	}
	return MakeHost(), false
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
