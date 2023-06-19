package binary_utils

import (
	"github.com/charmbracelet/log"
	"io"
	"io/ioutil"
	"os"
)

func ReadBinaryFile(filename string) ([]byte, int) {

	content, _ := ioutil.ReadFile(filename)

	fileSizeWithMb := float64(len(content)) / 1024.0 / 1024.0
	log.Debugf("Reading file: %s length is %.2fM -> %d %d bytes", filename, fileSizeWithMb, len(content))

	return content, len(content)
}

func AppendBinaryFile(filename string, data []byte) {
	fp, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fp.Close()

	fp.Seek(0, io.SeekEnd)
	fp.Write(data)
}

func WriteBinaryFile(filename string, data []byte) {
	fp, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	defer fp.Close()

	fp.Write(data)
}
