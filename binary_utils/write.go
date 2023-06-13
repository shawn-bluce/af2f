package binary_utils

import (
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"io/ioutil"
	"os"
)

func WriteBinaryFile(filename string, data []byte) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	bigFileContent, err := ioutil.ReadFile(filename)
	bigFileSize := len(bigFileContent)
	fileSize := float32(len(bigFileContent)) / 1024.0 / 1024.0
	log.Debugf("file: %s length is %.2fM", filename, fileSize)

	fp.Seek(int64(bigFileSize), 1)
	fp.Write(data)

	buff := make([]byte, 55) // 55=该文本的长度

	for {
		lens, err := fp.Read(buff)
		if err == io.EOF || lens < 0 {
			break
		}
	}
	fmt.Println("--------")
	fmt.Print(string(buff))
	fmt.Println("--------")

}
