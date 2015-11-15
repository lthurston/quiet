package host_test

import (
	"testing"

	"github.com/lthurston/quiet/host"
)

func TestParse0(t *testing.T) {
	hosts := host.HostsCollection{}
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
	hosts := host.HostsCollection{}
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
	if hosts.GetIndex(0).StartLine() != 1 {
		t.Errorf("Expecting first host to begin on line 1; got %v", hosts.GetIndex(0).StartLine())
	}
	if hosts.GetIndex(0).EndLine() != 7 {
		t.Errorf("Expecting first host to end on line 7; got %v", hosts.GetIndex(0).EndLine())
	}
}

func TestParse2(t *testing.T) {
	hosts := host.HostsCollection{}
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

#2A
#1A
#0A
Host fraceass
        Hostname fesces.org
        User blah blah woof woof
        Identidfysdfjkh dsf
        IdentityFile test
        IdentityFile blah blah
#0B
#1B
#2B

# A comment with some empty lines around it

`)
	if count := hosts.Count(); count != 3 {
		t.Errorf("Expecting 3 hosts in this config; got %v", count)
	}
	if hosts.GetIndex(0).StartLine() != 12 {
		t.Errorf("Expecting first host to begin on line 12; got %v", hosts.GetIndex(0).StartLine())
	}
	if hosts.GetIndex(0).EndLine() != 15 {
		t.Errorf("Expecting first host to end on line 15; got %v", hosts.GetIndex(0).EndLine())
	}
	if hosts.GetIndex(1).StartLine() != 17 {
		t.Errorf("Expecting second host to begin on line 17; got %v", hosts.GetIndex(1).StartLine())
	}
	if hosts.GetIndex(1).EndLine() != 20 {
		t.Errorf("Expecting second host to end on line 20; got %v", hosts.GetIndex(1).EndLine())
	}
	if hosts.GetIndex(2).StartLine() != 25 {
		t.Errorf("Expecting second host to begin on line 25; got %v", hosts.GetIndex(1).StartLine())
	}
	if hosts.GetIndex(2).EndLine() != 30 {
		t.Errorf("Expecting second host to end on line 30; got %v", hosts.GetIndex(1).EndLine())
	}
}

func TestFindHostByName(t *testing.T) {
	hosts := host.HostsCollection{}
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
