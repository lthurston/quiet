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

	if hosts.Hosts[0].GetName() != "skeins" {
		t.Errorf("Wrong host name, expecting: %s got: %s", name, hosts.Hosts[0].GetName())
	}

	if !reflect.DeepEqual(hosts.Hosts[0].GetAliases(), aliases) {
		t.Errorf("Wrong aliases, expecting: %s got: %s", aliases, hosts.Hosts[0].GetAliases())
	}

	if !reflect.DeepEqual(hosts.Hosts[0].GetConfig(), config) {
		t.Errorf("Wrong config, expecting: %s got: %s", config, hosts.Hosts[0].GetConfig())
	}
}
