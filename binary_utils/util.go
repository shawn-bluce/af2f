package binary_utils

import (
	"github.com/charmbracelet/log"
	"io/ioutil"
	"os"
)

func ReadBinaryFile(filename string) ([]byte, int) {

	content, _ := ioutil.ReadFile(filename)
	fileSize := float32(len(content)) / 1024.0 / 1024.0
	log.Debugf("file: %s length is %.2fM -> %d", filename, fileSize, len(content))

	return content, len(content)
}

func AppendBinaryFile(filename string, data []byte) {
	fileContent, _ := ioutil.ReadFile(filename)
	//bigFileSize := len(fileContent)
	fileSize := float32(len(fileContent)) / 1024.0 / 1024.0
	log.Debugf("file: %s length is %.2fM -> %d", filename, fileSize, len(fileContent))

	fp, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fp.Close()

	fp.Seek(0, 2)

	fp.Write(data)

}

func WriteBinaryFile(filename string, data []byte) {
	fp, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	defer fp.Close()

	fp.Write(data)
}
