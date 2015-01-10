package host

import (
	"bytes"
	"strings"
	"text/template"
)

// Host holds host info
type Host struct {
	name      string
	aliases   []string
	config    map[string]string
	startLine int
	endLine   int
}

// AddConfigFromString scans a line and adds a config item for it
func (host *Host) AddConfigFromString(line string) {
	line = strings.TrimSpace(line)
	sepIndex := strings.IndexAny(line, " \t")
	config, value := line[0:sepIndex], line[sepIndex+1:]
	host.config[config] = value
}

// SetName sets name
func (host *Host) SetName(name string) {
	host.name = name
}

// GetName gets name
func (host Host) GetName() string {
	return host.name
}

// SetAliases sets aliases
func (host *Host) SetAliases(aliases []string) {
	host.aliases = aliases
}

// GetAliases gets aliases
func (host Host) GetAliases() []string {
	return host.aliases
}

// SetConfig sets config
func (host *Host) SetConfig(config map[string]string) {
	host.config = config
}

// GetConfig gets config
func (host Host) GetConfig() map[string]string {
	return host.config
}

// SetStartLine sets start line
func (host *Host) SetStartLine(startLine int) {
	host.startLine = startLine
}

// GetStartLine get startLine
func (host Host) GetStartLine() int {
	return host.startLine
}

// SetEndLine sets the end line
func (host *Host) SetEndLine(endLine int) {
	host.endLine = endLine
}

// GetEndLine gets the end line
func (host Host) GetEndLine() int {
	return host.endLine
}

// RenderSnippet renders a host snippet
func (host Host) String() string {
	tmpl, err := template.New("snip").Parse(`
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

// MakeHost returns a Host!
func MakeHost() Host {
	i := Host{}
	i.SetConfig(make(map[string]string))
	return i
}
