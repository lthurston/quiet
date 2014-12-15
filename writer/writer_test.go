package writer_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/lthurston/quiet/writer"
)

var contents = `0
1
2
3
4
5
6
7
8
9`

var spliceContent = `I've
Been
Spliced
In`

// func TestGetLines(*testing.T) {
// 	r := strings.NewReader(contents)
// 	scanner := bufio.NewScanner(r)
// 	fmt.Println("TEST 1")
// 	fmt.Println(writer.GetLines(0, 4, scanner))
// 	r = strings.NewReader(contents)
// 	scanner = bufio.NewScanner(r)
// 	fmt.Println("TEST 2")
// 	fmt.Println(writer.GetLines(4, 0, scanner))
//
// 	r = strings.NewReader(contents)
// 	scanner = bufio.NewScanner(r)
// 	fmt.Println("TEST 3")
// 	fmt.Println(writer.GetLines(3, 7, scanner))
// }

func TestSpliceIntoMiddle(t *testing.T) {
	reader := strings.NewReader(contents)
	var buffer bytes.Buffer
	writer.SpliceInto(3, 8, spliceContent, reader, &buffer)
	splice := buffer.String()

	expectedResult := `0
1
2
I've
Been
Spliced
In
9`
	if expectedResult != splice {
		t.Error("Unexpected splice results")
	}
}

func TestSpliceIntoBeginning(t *testing.T) {
	reader := strings.NewReader(contents)
	var buffer bytes.Buffer
	writer.SpliceInto(0, 7, spliceContent, reader, &buffer)
	splice := buffer.String()

	expectedResult := `I've
Been
Spliced
In
8
9`
	if expectedResult != splice {
		t.Error("Unexpected splice results")
	}
}

func TestSpliceIntoEnd(t *testing.T) {
	reader := strings.NewReader(contents)
	var buffer bytes.Buffer
	writer.SpliceInto(8, 0, spliceContent, reader, &buffer)
	splice := buffer.String()

	expectedResult := `0
1
2
3
4
5
6
7
I've
Been
Spliced
In`
	if expectedResult != splice {
		fmt.Println("END")
		fmt.Println(splice)
		t.Error("Unexpected splice results")
	}
}
