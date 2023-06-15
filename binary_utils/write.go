package binary_utils

import (
	"fmt"
	"github.com/charmbracelet/log"
	"io/ioutil"
	"os"
)

func AppendBinaryFile(filename string, data []byte) {
	fileContent, _ := ioutil.ReadFile(filename)
	//bigFileSize := len(fileContent)
	fileSize := float32(len(fileContent)) / 1024.0 / 1024.0
	log.Debugf("file: %s length is %.2fM -> %d", filename, fileSize, len(fileContent))

	fp, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fp.Close()

	fp.Seek(0, 2)

	res, _ := fp.Write(data)

	fmt.Println("================")
	fmt.Println(res)
	fmt.Println("================")
}
