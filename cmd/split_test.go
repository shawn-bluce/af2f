package cmd

import (
	"os"
	"testing"
)

func preTest(filenameA, filenameB string) {

	err := os.Remove(filenameA)
	if err != nil {
		return
	}
	err = os.Remove(filenameB)
	if err != nil {
		return
	}
}

func TestValidateSplitArgs(t *testing.T) {
	filenameA := "filenameA"
	filenameB := "filenameB"
	preTest(filenameA, filenameB)

	if validateSplitArgs(filenameA, filenameB) {
		t.Errorf("filenameA and filenameB not exists")
	}

	os.Create(filenameA)
	if !validateSplitArgs(filenameA, filenameB) {
		t.Errorf("filenameA is exists")
	}

	os.Create(filenameB)
	if validateSplitArgs(filenameA, filenameB) {
		t.Errorf("filenameB is exists")
	}

	os.Remove(filenameA)
	os.Remove(filenameB)
}
