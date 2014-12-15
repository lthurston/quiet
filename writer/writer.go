package writer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// SpliceInto splices a new section into a buffer
func SpliceInto(lineStart, lineEnd int, spliceContent string, readSeeker io.ReadSeeker, buffer *bytes.Buffer) {
	var parts []string

	scanner := bufio.NewScanner(readSeeker)

	if lineStart > 0 {
		firstPart := getLines(0, lineStart, scanner)
		if len(firstPart) > 0 {
			parts = append(parts, firstPart)
		}
	}

	if len(spliceContent) > 0 {
		parts = append(parts, spliceContent)
	}

	if lineEnd > 0 {
		// Reset the reader so we can read again
		readSeeker.Seek(0, 0)
		scanner = bufio.NewScanner(readSeeker)
		lastPart := getLines(lineEnd, 0, scanner)
		if len(lastPart) > 0 {
			parts = append(parts, lastPart)
		}
	}

	combinedParts := strings.Join(parts, "\n")

	buffer.WriteString(combinedParts)
}

// GetLines gets a section of lines from lineStart to lineEnd
func getLines(lineStart, lineEnd int, scanner *bufio.Scanner) string {
	var lines []string
	lineIndex := 0

	// Skip the junk
	if lineStart > 0 {
		for scanner.Scan() {
			if lineStart == lineIndex {
				break
			}
			lineIndex++
		}
	}

	// Grab the goods
	for scanner.Scan() {
		if lineEnd == lineIndex {
			break
		}
		lines = append(lines, scanner.Text())
		lineIndex++
	}

	return strings.Join(lines, "\n")
}
