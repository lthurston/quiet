package host_test

import (
	"reflect"
	"testing"

	h "github.com/lthurston/quiet/host"
	"github.com/lthurston/quiet/parser"
)

func TestString(t *testing.T) {
	name := "skeins"
	aliases := []string{"sizzle", "toast", "blort"}
	config := map[string]string{"User": "gushaguy", "HostName": "skeins.com"}

	host := h.MakeHost()
	host.SetName(name)
	host.SetAliases(aliases)
	host.SetConfig(config)
	host.SetEndLine(0)
	host.SetStartLine(0)

	rendered := host.String()

	hosts := parser.HostsCollection{}
	hosts.ReadFromString(rendered)

	if hosts.GetIndex(0).Name() != "skeins" {
		t.Errorf("Wrong host name, expecting: %s got: %s", name, hosts.GetIndex(0).Name())
	}

	if !reflect.DeepEqual(hosts.GetIndex(0).Aliases(), aliases) {
		t.Errorf("Wrong aliases, expecting: %s got: %s", aliases, hosts.GetIndex(0).Aliases())
	}

	if !reflect.DeepEqual(hosts.GetIndex(0).Config(), config) {
		t.Errorf("Wrong config, expecting: %s got: %s", config, hosts.GetIndex(0).Config())
	}
}
