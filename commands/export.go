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
	"encoding/base64"
	"crypto/aes"
	"io"
	"crypto/rand"
	"crypto/cipher"
	"errors"
	"compress/gzip"
	"math"
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
				key := []byte("something something something  1")
				tarOut := tarOutput(archive)
				zippedTarOut := zip(tarOut)
				encrypted, _ := encrypt(key, zippedTarOut)
				base64d := base64.StdEncoding.EncodeToString(encrypted)

				out := ""
				for i := 1.00 ; i <= math.Ceil(float64(len(base64d)) / 80.0); i++ {
					if (i * 80) <= float64(len(base64d)) {
						out = out + base64d[int((i - 1) * 80):int(i * 80)]
					} else {
						out = out + base64d[int((i - 1) * 80):]
					}
				}

				backBase64d := strings.Join(out, "\n")
				backEncryped := base64.StdEncoding.DecodeString(backBase64d)
				backZippedTarOut, _ := decrypt(key, backEncrypted)
				backTarOut :=


				//decrypted, _ := decrypt(aKey, ecrypted)
				//fmt.Println(string(decrypted))
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



