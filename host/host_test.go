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

func TestContainsStrings(t *testing.T) {
	name := "skeins"
	aliases := []string{"sizzle", "toast", "blort"}
	options := []string{"User gushaguy", "HostName skeins.com"}
	host := h.MakeHost()
	host.SetName(name)
	host.SetAliases(aliases)

	for _, line := range options {
		host.AddOptionFromString(line)
	}

	if(!host.ContainsStrings([]string{"sizzle", "ast", "guy", ".com"})) {
		t.Errorf("Contains strings isn't returning true when it should")
	}

	if(host.ContainsStrings([]string{"poo"})) {
		t.Errorf("Contains strings isn't returning false when it should")
	}
}

func TestAddOptionFromString(t *testing.T) {
	host := h.MakeHost()

	host.AddOptionFromString("junktown ~/.sdklfj/hosrsefj")
	host.AddOptionFromString("Mule")
	host.AddOptionFromString(" 	sdlkfj 	 	 	 ")
	host.AddOptionFromString(" 	sddfsdflkfj 	 sdfsdf	 	sdfsdfsdf ")


	if arg := host.GetOptionArgument("junktown"); arg != "~/.sdklfj/hosrsefj" {
		t.Errorf("Unexpected option argument. Expected \"~/.sdklfj/hosrsefj\" got %s.", arg)
	}

	if arg := host.GetOptionArgument("Mule"); arg != "" {
		t.Errorf("Unexpected option argument. Expected empty string got %s.", arg)
	}

	if arg := host.GetOptionArgument("sdlkfj"); arg != "" {
		t.Errorf("Unexpected option argument. Expected empty string got \"%s\".", arg)
	}

	if arg := host.GetOptionArgument("sddfsdflkfj"); arg != "sdfsdf	 	sdfsdfsdf" {
		t.Errorf("Unexpected option argument. Expected \"sdfsdf	 	sdfsdfsdf\" got \"%s\".", arg)
	}
}