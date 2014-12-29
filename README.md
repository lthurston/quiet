# Quiet

Quiet simplifies repetitive SSH config management tasks. Allows listing,
editing, copying, exporting (to share with others).

## Usage

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
  new                       Appends new host using other host as template
  help [command]            Help about any command

 Available Flags:
      --help=false: help for quiet

Use "quiet help [command]" for more information about that command.
```

## Installation

1) Install Go (brew install go, or whatever)

2) Set up a go dev environment: `mkdir -p ~/go/{src,bin}`

3) Add some stuff to your .bash_profile:
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

4) Get theeself to the Go-ery: `cd ~/go`

5) Fetch the source: `go get github.com/lthurston/quiet`

6) Compile and install: `go install src/github.com/lthurston/quiet/quiet.go`
