package encrypt_tool

import (
	"crypto/aes"
	"crypto/md5"
	"fmt"
	"github.com/charmbracelet/log"
)

func getPasswordHash(algorithm string, password string) string {
	passwordByBytes := []byte(password)
	passwordWithHash := md5.Sum(passwordByBytes)
	hexPassword := fmt.Sprintf("%x", passwordWithHash) //将[]byte转成16进制

	switch algorithm {
	default:
		return "NOT_FOUND" // should never run this
	case "aes-128":
		return hexPassword[:16]
	case "aes-192":
		return hexPassword[:24]
	case "aes-256":
		return hexPassword[:32]
	}
}

func Encrypt(data []byte, algorithm string, password string) []byte {
	passwordByHash := getPasswordHash(algorithm, password)
	cipher, err := aes.NewCipher([]byte(passwordByHash))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	plaintext := []byte(data)
	ciphertext := make([]byte, len(plaintext))
	cipher.Encrypt(ciphertext, plaintext)
	log.Infof("Encrypted by password %s -> %s", password, passwordByHash)
	log.Infof("plaintext first 32bytes: %x", plaintext[:32])
	log.Infof("ciphertext first 32bytes: %x", ciphertext[:32])
	return ciphertext
}

func Decrypt(data []byte, algorithm string, password string) []byte {
	passwordByHash := getPasswordHash(algorithm, password)
	cipher, err := aes.NewCipher([]byte(passwordByHash))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	plaintext := []byte(data)
	ciphertext := make([]byte, len(plaintext))
	cipher.Decrypt(ciphertext, plaintext)
	log.Infof("Decrypted by password %s -> %s", password, passwordByHash)
	log.Infof("ciphertext first 32bytes: %x", ciphertext[:32])
	log.Infof("plaintext first 32bytes: %x", plaintext[:32])
	return plaintext
}
