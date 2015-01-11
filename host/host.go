package host

import (
	"bytes"
	"strings"
	"text/template"
)

// Host holds host info
type Host struct {
	name    string
	aliases []string
	config  map[string]string
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

// SetConfig sets config
func (host *Host) SetConfig(config map[string]string) {
	host.config = config
}

// Config gets config
func (host Host) Config() map[string]string {
	return host.config
}

// RenderSnippet renders a host snippet
func (host Host) String() string {
	tmpl, err := template.New("snip").Parse(`
Host {{.Name}}{{if .Aliases}}{{range .Aliases}} {{.}}{{end}}{{end}}
  {{range $key, $value := .Config	 }}{{$key}} {{$value}}
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
