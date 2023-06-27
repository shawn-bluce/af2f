package cmd

import (
	"os"
	"testing"
)

func TestValidateAppendArgs(t *testing.T) {
	var validated bool
	filenameA := "DELETE_ME_A"
	filenameB := "DELETE_ME_B"
	password := "123456"
	algorithm := "aes-128"

	preTest(filenameA, filenameB)

	validated = validateAppendArgs(filenameA, filenameB, password, algorithm)
	if validated {
		t.Errorf("filenameA and filenameB exists")
	}

	os.Create(filenameA)

	validated = validateAppendArgs(filenameA, filenameB, password, algorithm)
	if validated {
		t.Errorf("filenameB exists")
	}

	os.Create(filenameB)
	validated = validateAppendArgs(filenameA, filenameB, password, algorithm)
	if !validated {
		t.Errorf("this args should be pass")
	}

	validated = validateAppendArgs(filenameA, filenameB, "12345", algorithm)
	if validated {
		t.Errorf("password is 12345, length < 6, shoudent pass")
	}

	passwordMoreThan32 := "012345678901234567890123456789ABC"
	validated = validateAppendArgs(filenameA, filenameB, passwordMoreThan32, algorithm)
	if validated {
		t.Errorf("password length > 32, shoudent pass")
	}

	validated = validateAppendArgs(filenameA, filenameB, password, "aes-233")
	if validated {
		t.Errorf("algorithm is aes-233, shoudent pass")
	}

	os.Remove(filenameA)
	os.Remove(filenameB)
}

func TestDoAppend(t *testing.T) {
	filenameA := "filenameA"
	filenameB := "filenameB"

	if fileNotExists(filenameA) {
		os.Create(filenameA)
	}
	if fileNotExists(filenameB) {
		os.Create(filenameB)
	}

	doAppend(filenameA, filenameB, "aes-128", "1223456")
	doAppend(filenameA, filenameB, "aes-192", "1223456")
	doAppend(filenameA, filenameB, "aes-256", "1223456")
	doAppend(filenameA, filenameB, "aes-256", "")
}
