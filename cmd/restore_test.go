package cmd

import (
	"os"
	"testing"
)

func TestValidateRestoreArgs(t *testing.T) {
	filename := "DELETE_ME"
	if !fileNotExists(filename) {
		os.Remove(filename)
	}
	if validateRestoreArgs(filename) {
		t.Errorf("file should not exists")
	}

	os.Create(filename)
	if !validateRestoreArgs(filename) {
		t.Errorf("file should exists")
	}

	os.Remove(filename)
}

func TestDoRestore(t *testing.T) {
	filenameA := "filenameA"
	filenameB := "filenameB"

	if fileNotExists(filenameA) {
		os.Create(filenameA)
	}
	if fileNotExists(filenameB) {
		os.Create(filenameB)
	}

	doAppend(filenameA, filenameB, "aes-128", "123456")
	doRestore(filenameA)

	if !fileNotExists(filenameA) {
		os.Remove(filenameA)
	}
	if !fileNotExists(filenameB) {
		os.Remove(filenameB)
	}
}
