package packager

import (
	"fmt"
	"bytes"
	"archive/tar"
	"log"
)

func hrOutput(archive exportArchive) {
	fmt.Println(archive.config)

	for _, file := range archive.files {
		fmt.Println(file.filename, "\n")
		fmt.Println(string(file.contents))
	}
}

func tarOutput(archive exportArchive) []byte {
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
	return buf.Bytes()
}

func tarInput([]byte) archive exportArchive {


}