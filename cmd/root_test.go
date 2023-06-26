package cmd

import (
	"os"
	"testing"
)

func TestFileNotExists(t *testing.T) {
	filename := "DELETE_ME"
	if !fileNotExists(filename) {
		t.Errorf("filename: DELETE_ME shouldent exists")
	}

	os.Create(filename)
	if fileNotExists(filename) {
		t.Errorf("filename: DELETE_ME should exists")
	}

	os.Remove(filename)
	if !fileNotExists(filename) {
		t.Errorf("filename: DELETE_ME shouldent exists")
	}
}
