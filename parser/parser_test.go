package parser_test

import (
	"testing"

	"github.com/lthurston/quiet/parser"
)

func TestParse0(t *testing.T) {
	hosts := parser.HostsCollection{}
	hosts.ReadFromString(
		`
Nothing in here!

# Something something something
Host *
Something garbagey

# Host Crap

`)
	if count := hosts.Count(); count != 0 {
		t.Errorf("Expecting 0 hosts in this config; got %v", count)
	}
}

func TestParse1(t *testing.T) {
	hosts := parser.HostsCollection{}
	hosts.ReadFromString(
		`
Host junkfood
	Hostname junkfood.com
	User cavityman

	IndentityFile ~/.ssh/id_rsa

	Something Else in here

# A comment


`)
	if count := hosts.Count(); count != 1 {
		t.Errorf("Expecting 1 hosts in this config; got %v", count)
	}
	if hosts.GetIndex(0).StartLine() != 2 {
		t.Errorf("Expecting first host to begin on line 2; got %v", hosts.GetIndex(0).StartLine())
	}
	if hosts.GetIndex(0).EndLine() != 8 {
		t.Errorf("Expecting first host to end on line 8; got %v", hosts.GetIndex(0).EndLine())
	}
}

func TestParse2(t *testing.T) {
	hosts := parser.HostsCollection{}
	hosts.ReadFromString(
		`
Host *
	IdentitiesOnly yes

# That one shouldn't become a host
# because it is "*"


#################################################
# Ormulex - id_rsa ps is Billy's usual
#################################################

Host ormulex-dev
			Hostname dev.ormulex.com
			User mule
			IdentityFile ~/.ssh/ormulex/id_rsa

Host ormulex-qa
			Hostname cmw3wbq10
			User monkey
			ProxyCommand ssh cmw3wbq10 'nc %h %banana %p'


# A comment with some empty lines around it

`)
	if count := hosts.Count(); count != 2 {
		t.Errorf("Expecting 2 hosts in this config; got %v", count)
	}
	if hosts.GetIndex(0).StartLine() != 13 {
		t.Errorf("Expecting first host to begin on line 13; got %v", hosts.GetIndex(0).StartLine())
	}
	if hosts.GetIndex(0).EndLine() != 16 {
		t.Errorf("Expecting first host to end on line 16; got %v", hosts.GetIndex(0).EndLine())
	}
	if hosts.GetIndex(1).StartLine() != 18 {
		t.Errorf("Expecting second host to begin on line 18; got %v", hosts.GetIndex(1).StartLine())
	}
	if hosts.GetIndex(1).EndLine() != 21 {
		t.Errorf("Expecting second host to end on line 21; got %v", hosts.GetIndex(1).EndLine())
	}
}

func TestFindHostByName(t *testing.T) {
	hosts := parser.HostsCollection{}
	hosts.ReadFromString(
		`
Host *
		IdentitiesOnly yes

# That one shouldn't become a host
# because it is "*"


#################################################
# Ormulex - id_rsa ps is Billy's usual
#################################################

Host ormulex-dev
				Hostname dev.ormulex.com
				User mule
				IdentityFile ~/.ssh/ormulex/id_rsa

Host ormulex-qa
				Hostname cmw3wbq10
				User monkey
				ProxyCommand ssh cmw3wbq10 'nc %h %banana %p'


# A comment with some empty lines around it

`)

	host, found := hosts.FindHostByName("ormulex-dev")
	if !found {
		t.Errorf("Couldn't find the host using hosts.FindHostByName! (#1)")
	}
	if found && host.Name() != "ormulex-dev" {
		t.Errorf("hosts.FindHostByName found the wrong host (#2)")
	}
	host, found = hosts.FindHostByName("ormulex-qa")
	if !found {
		t.Errorf("Couldn't find the host using hosts.FindHostByName! (#3)")
	}
	if found && host.Name() != "ormulex-qa" {
		t.Errorf("hosts.FindHostByName found the wrong host (#4)")
	}

}

func TestFindHostValue(t *testing.T) {

}

func TestWriteNewHost(t *testing.T) {

}
