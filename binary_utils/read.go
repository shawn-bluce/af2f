package binary_utils

import (
	"github.com/charmbracelet/log"
	"io/ioutil"
)

func ReadBinaryFile(filename string) []byte {

	content, _ := ioutil.ReadFile(filename)
	fileSize := float32(len(content)) / 1024.0 / 1024.0
	log.Debugf("file: %s length is %.2fM", filename, fileSize)

	return content
}
