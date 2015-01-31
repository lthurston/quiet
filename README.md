# Quiet

Quiet simplifies repetitive SSH config management tasks. Allows listing,
editing, copying, exporting (to share with others).

## Usage

### General Usage

Help output says most of it:

```

Usage:
  quiet [flags]
  quiet [command]

Available Commands:
  version                   Print the version number of Quiet
  config                    Lists / modifies configuration
  dump                      Dumps everything Quiet knows
  list                      Lists all hosts
  export                    Exports config snippet and keys for sharing
  undo                      Undo the last quiet action
  warp                      Generates warp config
  rm                        Deletes a host
  new                       Appends new host using other host as template
  help [command]            Help about any command

 Available Flags:
      --help=false: help for quiet

Use "quiet help [command]" for more information about that command.
```

### New Command

New is currently the most useful command. It will create a new host snippet
from an existing snippet and allow manipulation of each of the config values (
host aliases currently excluded). Here's the output of `quiet help new`:

```
New appends a new host to your SSH configuration based on other host

Usage:
quiet new [flags]

Available Flags:
-f, --from="": host to use as template
--help=false: help for new
-n, --name="": new host name
-s, --skip-interactive=false: just copy; don't allow interactive
-o, --stdout=false: output to stdout rather than appending SSH config file


Use "quiet help [command]" for more information about that command.
```

The -f/--from setting can also come from .quiet configuration as new.from. The
value should be an existing hostname or a template that you've created for
purpose of providing common defaults.

-s/--skip-interactive will skip the interactive modification of values from the
copied configuration snippet, making new behave more like a copy action.

-n/--name allows the name of the new host to be specified in the command.

-o/--stdout will output the new snippet rather than writing to the ssh config
file.

**Examples**

_You want to copy an existing configuration snippet, but modify one or more
of the host configuration lines:_

```
$ quiet new -foldhostname
Copying host "oldhostname"
Name: newhostname
Hostname [default is "oldhostname.com"]: newhostname.com
User [default is "frank"]: ubuntu
IdentityFile [default is "~/.ssh/id_oldhostname"]: ~/.ssh/id_ubuntu
```

_You want to copy your default new configuration snippet, and keep everything
but the hostname the same:_

In your .quiet config, make sure that new.from is set to your default host to
use as a template.

```
[new]
from = "ubuntu"
```

Then invoke quiet like this:

```
quiet new -nnewhost -s

Copying host "ubuntu"
```

Here's what quiet added to your SSH configuration:

```
Host newhost
  Hostname test.com
  IdentityFile ~/.ssh/id_default
  User ubuntu

```

Note that the hostname wasn't updated. I might add a way to do that from the
command line, perhaps by repurposing the -n/--name as to the hostname value.

Also, you can use the -o flag as a kind of dry run to see what quiet will append
to your hosts file before you have it actually modify the file.

### Backups

When quiet updates your config file it backs up the most recent version of the
file, and keeps the last 5 versions. The backup file has the same name as the
SSH configuration file with `.quiet.bak.#` appended to the filename, where the
octothrombulus is replaced by a numbrosimbel (1 - 5).

You can revert to the last backup with `quiet undo`. This just swaps the most
recent backup with the current file.

I'm not gonna promise Quiet won't mangle your config file in some sort of permanent
way, especially while it's still alpha, so maybe you should just back up your own
stuff in your own way before you start using it. I'm only a human being.

## Installation

1) Install Go (brew install go, or whatever)

2) Set up a go dev environment: `mkdir -p ~/go/{src,bin}`

3) Add some stuff to your .bash_profile:
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

4) Get thee to the Go-ery: `cd ~/go`

5) Fetch the source: `go get github.com/lthurston/quiet`

6) Compile and install: `go install src/github.com/lthurston/quiet/quiet.go`

### Release Notes

0.3 alpha

* Updated this readme
* rm command implemented
* dump command implemented
* refactoring of the parser / hosts collection

0.2 alpha

* refactored hosts / host positions

0.1 alpha

* made it work