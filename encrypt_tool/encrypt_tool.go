package encrypt_tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
)

func pkcs7Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize
	paddingData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, paddingData...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	paddingLength := int(origData[length-1])
	return origData[:length-paddingLength]
}

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

func AESEncrypt(clearText []byte, algorithm string, password string) []byte {
	passwordWithByte := []byte(getPasswordHash(algorithm, password))

	block, _ := aes.NewCipher(passwordWithByte)
	clearText = pkcs7Padding(clearText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, passwordWithByte[:block.BlockSize()])
	cipherText := make([]byte, len(clearText))
	blockMode.CryptBlocks(cipherText, clearText)

	return cipherText
}

func AESDecrypt(cipherText []byte, algorithm string, password string) []byte {
	passwordWithByte := []byte(getPasswordHash(algorithm, password))

	block, _ := aes.NewCipher(passwordWithByte)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, passwordWithByte[:blockSize])
	clearText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(clearText, cipherText)
	clearText = pkcs7UnPadding(clearText)
	return clearText
}
