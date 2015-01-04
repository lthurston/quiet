package writer

import (
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/lthurston/quiet/config"
)

type fileNameMap struct {
	from string
	to   string
}

const (
	appendGlob = ".quiet.bak.*"
	tempGlob   = ".quiet.tmp.*"
)

// Undo swaps most recent backup with previous
func Undo() error {
	var err error
	filename := config.GetConfigFile()
	firstBackup := replaceGlobWithInt(filename+appendGlob, 1)
	tmpFilename := getTempFilename(filename + tempGlob)

	// Not handling errors at all -- fix it
	renameFile(firstBackup, tmpFilename)
	renameFile(filename, firstBackup)
	renameFile(tmpFilename, filename)

	return err
}

func getTempFilename(filename string) string {
	exists := true
	var filenameR string
	for exists {
		filenameR = replaceGlobWithInt(filename, rand.Int())
		exists, _ = fileExists(filenameR)
	}
	return filenameR
}

// Replace replaces the file
func Replace(content string) error {
	err := backup()
	return err
}

// Append does what it sounds like
func Append(contents string) error {
	f, err := os.OpenFile(config.GetConfigFile(), os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	err = backup()
	if err != nil {
		return err
	}

	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}
	return err
}

func backup() error {
	filename := config.GetConfigFile()
	backupsToStore := 5
	err := rotateBackups(getBackupRotationMapping(backupsToStore, filename+appendGlob))
	if err != nil {
		return err
	}
	firstBackup := replaceGlobWithInt(filename+appendGlob, 1)
	err = copyFile(filename, firstBackup)
	if err != nil {
		return err
	}

	return err
}

func rotateBackups(rotationMapping []fileNameMap) error {
	var err error
	for _, renameMappings := range rotationMapping {
		exists, err := fileExists(renameMappings.to)
		if err != nil {
			return err
		}
		if exists {
			err := os.Remove(renameMappings.to)
			if err != nil {
				return err
			}
		}
		err = renameFile(renameMappings.from, renameMappings.to)
		if err != nil {
			return err
		}
	}
	return err
}

func copyFile(source, target string) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	err = out.Close()
	return err
}

func renameFile(source, target string) error {
	var exists bool
	var err error
	exists, err = fileExists(source)
	if err != nil {
		return err
	}
	if exists {
		err = os.Rename(source, target)
		if err != nil {
			return err
		}
	}
	return err
}

func fileExists(filename string) (bool, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	var err2 error
	return fi != nil, err2
}

func getBackupRotationMapping(count int, appendGlob string) (rotationMapping []fileNameMap) {
	for i := count - 1; i >= 1; i-- {
		k := replaceGlobWithInt(appendGlob, i)
		v := replaceGlobWithInt(appendGlob, i+1)
		rotationMapping = append(rotationMapping, fileNameMap{from: k, to: v})
	}
	return
}

func replaceGlobWithInt(appendGlob string, i int) string {
	return strings.Replace(appendGlob, "*", strconv.Itoa(i), 1)
}
