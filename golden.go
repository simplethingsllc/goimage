package gotransform

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func writeToFile(buf []byte, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(buf)
	f.Sync()
}

// GoldenCompare checks the given buffer against the raw filename data. If the file doesn't exist,
// then the buffer is written out to the file.
func GoldenCompare(t *testing.T, buf []byte, filename string) {
	if !fileExists(filename) {
		writeToFile(buf, filename)
		return
	}
	golden, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if bytes.Compare(buf, golden) != 0 {
		t.Fatalf("Golden file not equal to given")
	}
}
