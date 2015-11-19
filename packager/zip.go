package packager
import (
	"compress/gzip"
	"bytes"
)

func zip(contents []byte) []byte {
	buf := &bytes.Buffer{}
	gw := gzip.NewWriter(buf)

	gw.Write(contents)
	gw.Close()

	return buf.Bytes()
}

func unzip(contents []byte) []byte {
	buf := &bytes.Buffer{}
	gw := gzip.NewWriter(buf)

	gw.Write(contents)
	gw.Close()

	return buf.Bytes()
}