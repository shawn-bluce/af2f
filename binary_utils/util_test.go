package binary_utils

import (
	"bytes"
	"os"
	"testing"
)

func preTest(filename string) {
	err := os.Remove(filename)
	if err != nil {
		return
	}
}

func TestWriteBinaryFile(t *testing.T) {
	filename := "test_filename"
	originData := []byte("hello, world")
	appendData := []byte("world, hello")

	preTest(filename)

	WriteBinaryFile(filename, originData)
	content, length := ReadBinaryFile(filename)

	if !bytes.Equal(originData, content) {
		t.Errorf("Write and Read file ERROR")
	}

	if length != 12 {
		t.Errorf("read file length ERROR")
	}

	resultShouldBeBytes := []byte("hello, worldworld, hello")
	AppendBinaryFile(filename, appendData)
	content, length = ReadBinaryFile(filename)

	if !bytes.Equal(resultShouldBeBytes, content) {
		t.Errorf("Write and Read file ERROR")
	}

	if length != 24 {
		t.Errorf("read file length ERROR")
	}

	os.Remove(filename)
}
