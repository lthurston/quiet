package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lthurston/quiet/host"
	"github.com/lthurston/quiet/config"
	"github.com/lthurston/quiet/util"
	"strings"
	"io/ioutil"
	"log"
	"archive/tar"
	"bytes"
)

var exportTar, exportIgnoreProblems bool

type exportArchiveFile struct {
	filename string
	contents []byte
}

type exportArchive struct {
	config string
	files []exportArchiveFile
}

var exportCmd = &cobra.Command{
	Use:   "export [hostname]",
	Short: "Exports config snippet and keys for sharing",
	Long:  `Exports one or more host snippets and associated keys for easy distribution`,
}

func init() {
	exportCmd.Flags().BoolVarP(&exportTar, "tar", "t", false, "output a tarball archive")
	exportCmd.Flags().BoolVarP(&exportIgnoreProblems, "who-cares", "w", false, "prevents fatal error for fixable problems")
	exportCmd.Run = exportCmdRun
}

func exportCmdRun(cmd *cobra.Command, args []string) {
	hosts := host.HostsCollection{}
	hosts.ReadFromFile(config.GetConfigFile())

	if len(args) > 0 {
		if host, found := hosts.FindHostByName(args[0]); found {
			archive := makeExportArchive(host)

			if (exportTar) {
				fmt.Println(tarOutput(archive))
			} else {
				hrOutput(archive)
			}
		} else {
			log.Fatal("Couldn't find host or empty: \"" + args[0] + "\"")
		}
	} else {
		log.Fatal("Hostname to export must be specified")
	}
}

func makeExportArchive(host host.HostPosition) exportArchive {
	archive := exportArchive{}
	archive.config = host.String()
	fields := strings.Split(config.GetConfigExportFilenameOptions(), ",")
	for _, field := range fields {
		filename := host.GetOptionArgument(strings.TrimSpace(field))
		if (len(filename) > 0) {
			if fileContents, err := ioutil.ReadFile(util.Detilde(filename)); err == nil {
				archive.files = append(archive.files, exportArchiveFile{filename, fileContents})

				// If there's a file by the same name with .pub at the end, show it too
				filenamePub := filename + ".pub"
				if fileContentsPub, err := ioutil.ReadFile(util.Detilde(filenamePub)); err == nil {
					archive.files = append(archive.files, exportArchiveFile{filenamePub, fileContentsPub})
				} else {
					ignorableFatal(err)
				}
			} else {
				ignorableFatal(err)
			}
		}
	}
	return archive
}

func ignorableFatal(err error) {
	if(!exportIgnoreProblems) {
		log.Fatal(err,"\nNote: this problem can be ignored with the -w / --who-cares flag")
	}
}

func hrOutput(archive exportArchive) {
	fmt.Println(archive.config)

	for _, file := range archive.files {
		fmt.Println(file.filename, "\n")
		fmt.Println(string(file.contents))
	}
}

func tarOutput(archive exportArchive) string {
	buf := &bytes.Buffer{}
	tw := tar.NewWriter(buf)

	// Note that because we don't actually expect this tarball to be extracted, we're not dealing with
	// setting modes or file times
	header := &tar.Header {
		Name: "quiet.ssh-config-snippet",
		Size: int64(len(archive.config)),
	}
	if err := tw.WriteHeader(header); err != nil {
		log.Fatal(err)
	}
	if _, err := tw.Write([]byte(archive.config)); err != nil {
		log.Fatal(err)
	}

	for _, file := range archive.files {
		header = &tar.Header {
			Name: file.filename,
			Size: int64(len(file.contents)),
		}
		if err := tw.WriteHeader(header); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write(file.contents); err != nil {
			log.Fatal(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}