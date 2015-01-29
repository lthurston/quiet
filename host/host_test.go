package host_test

import (
	"reflect"
	"testing"
	h "github.com/lthurston/quiet/host"
)

func TestString(t *testing.T) {
	name := "skeins"
	aliases := []string{"sizzle", "toast", "blort"}
	options := []string{"User gushaguy", "HostName skeins.com"}

	host := h.MakeHost()
	host.SetName(name)
	host.SetAliases(aliases)

	for _, line := range options {
		host.AddOptionFromString(line)
	}

	rendered := host.String()

	hosts := h.HostsCollection{}
	hosts.ReadFromString(rendered)

	if hosts.GetIndex(0).Name() != "skeins" {
		t.Errorf("Wrong host name, expecting: %s got: %s", name, hosts.GetIndex(0).Name())
	}

	if !reflect.DeepEqual(hosts.GetIndex(0).Aliases(), aliases) {
		t.Errorf("Wrong aliases, expecting: %s got: %s", aliases, hosts.GetIndex(0).Aliases())
	}

	// Need to rethink this comparison if option isn't exported
	// if !reflect.DeepEqual(hosts.GetIndex(0).Options(), config) {
	// 	t.Errorf("Wrong config, expecting: %s got: %s", config, hosts.GetIndex(0).Config())
	// }
}
