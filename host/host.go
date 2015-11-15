package host

import (
	"bytes"
	"strings"
	"text/template"
	"regexp"
)

// Option contains a keyword and an argument
type Option struct {
	keyword  string
	argument string
}

// MakeOption creates an optin
func MakeOption(keyword string, argument string) Option {
	return Option{keyword: keyword, argument: argument}
}

// Argument gets the argument
func (option Option) Argument() string {
	return option.argument
}

// Keyword gets the keyword
func (option Option) Keyword() string {
	return option.keyword
}

func (option Option) String() string {
	return option.keyword + " " + option.argument
}

// Host holds host info
type Host struct {
	name    string
	aliases []string
	options []Option
}

// GetOptionArgument gets an argument for an option keyword for a host
func (host Host) GetOptionArgument(keyword string) string {
	for _, option := range host.options {
		if option.keyword == keyword {
			return option.argument
		}
	}
	return ""
}

// AddOptionFromString scans a line and adds a config item for it
func (host *Host) AddOptionFromString(line string) {
	line = strings.TrimSpace(line)
	sepIndex := strings.IndexAny(line, " \t")

	var keyword, argument string
	if sepIndex != -1 {
		keyword, argument = line[0:sepIndex], line[sepIndex+1:]
	} else {
		keyword, argument = line, ""
	}

	host.options = append(host.options, Option{keyword: strings.TrimSpace(keyword), argument: strings.TrimSpace(argument)})
}

// SetName sets name
func (host *Host) SetName(name string) {
	host.name = name
}

// Name gets name
func (host Host) Name() string {
	return host.name
}

// SetAliases sets aliases
func (host *Host) SetAliases(aliases []string) {
	host.aliases = aliases
}

// Aliases gets aliases
func (host Host) Aliases() []string {
	return host.aliases
}

// SetOptions sets config
func (host *Host) SetOptions(options []Option) {
	host.options = options
}

// AddOption adds an option
func (host *Host) AddOption(keyword, argument string) {
	host.options = append(host.options, Option{keyword: keyword, argument: argument})
}

// Options gets config
func (host Host) Options() []Option {
	return host.options
}

// RenderSnippet renders a host snippet
func (host Host) String() string {
	tmpl, err := template.New("snip").Parse(`
Host {{.Name}}{{if .Aliases}}{{range .Aliases}} {{.}}{{end}}{{end}}
{{range .Options }}	{{.}}
{{end}}`)
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

// ContainsStrings returns true if the host contains the findStrings
func (host Host) ContainsStrings(findStrings []string) bool {
	for _, findString := range findStrings {
		escArg := regexp.QuoteMeta(findString)
		re := regexp.MustCompile("(?i)(" + escArg +")")
		if(!re.Match([]byte(host.String()))) {
			return false
		}
	}
	return true
}

// MakeHost returns a Host!
func MakeHost() Host {
	i := Host{}
	i.SetOptions([]Option{})
	return i
}
