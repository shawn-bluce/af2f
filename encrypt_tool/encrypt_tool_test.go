package encrypt_tool

import (
	"bytes"
	"testing"
)

func TestAESEncryptAndDecrypt(t *testing.T) {
	cleanData := []byte("hello, world")
	password := "correct password"

	algorithmList := []string{"aes-128", "aes-192", "aes-256"}

	for _, algorithm := range algorithmList {
		encryptedData := AESEncrypt(cleanData, algorithm, password)
		decryptedData := AESDecrypt(encryptedData, algorithm, password)
		if bytes.Equal(cleanData, decryptedData) == false {
			t.Errorf("same algorithm and password, but decrypt failed")
		} else {
			t.Logf("encrypt/decrypt with %s by %s success", algorithm, password)
		}
	}
}

func TestPasswordHash(t *testing.T) {
	if getPasswordHash("aes-233", "xx") == "NOT_FOUND" {
		t.Logf("have no aes-233 algorithm")
	} else {
		t.Errorf("don't should find algorithm aes-233")
	}
}
