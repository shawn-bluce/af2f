package cmd

import (
	"af2f/binary_utils"
	"af2f/common_utils"
	"encoding/binary"
	"github.com/charmbracelet/log"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func validateSplitArgs(file string, output string, password string, algorithm string) bool {
	validated := true
	if fileNotExists(file) {
		log.Errorf("-f file: %s is not exists", file)
		validated = false
	}
	if !fileNotExists(output) {
		log.Errorf("-o file: %s is already exists", output)
		validated = false
	}
	if len(password) > 32 {
		log.Errorf("password lens must less than 32")
		validated = false
	}

	algorithmIdFound := false
	for _, id := range common_utils.GetAlgorithmMap() {
		if algorithm == string(rune(id)) {
			algorithmIdFound = true
			break
		}
	}
	if !algorithmIdFound {
		log.Errorf("algorithm id: %s are not support", algorithm)
		validated = false
	}

	return validated
}

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a file from a file build from af2f",
	Run: func(cmd *cobra.Command, args []string) {
		bigFile, _ := cmd.Flags().GetString("file")
		outputFile, _ := cmd.Flags().GetString("output")
		encryption, _ := cmd.Flags().GetString("encryption")
		password, _ := cmd.Flags().GetString("password")
		if !validateSplitArgs(bigFile, outputFile, password, encryption) {
			log.Errorf("Do not pass the params validate")
			os.Exit(1)
		}

		fp, _ := os.OpenFile(bigFile, os.O_RDONLY, 0644)
		defer fp.Close()

		_, sumSize := binary_utils.ReadBinaryFile(bigFile)

		// read version
		fp.Seek(-8, io.SeekEnd)
		buffer := make([]byte, 8)
		fp.Read(buffer)
		version := binary.LittleEndian.Uint64(buffer)
		log.Debugf("read version is: %d", version)

		// read sourceBigFileSize
		fp.Seek(-16, io.SeekEnd)
		buffer = make([]byte, 8)
		fp.Read(buffer)
		sourceBigFileSize := binary.LittleEndian.Uint64(buffer)
		log.Debugf("read sourceBigFileSize is: %d", sourceBigFileSize)

		// read encryption algorithm
		fp.Seek(-24, io.SeekEnd)
		buffer = make([]byte, 8)
		fp.Read(buffer)
		algorithmId := binary.LittleEndian.Uint64(buffer)
		log.Debugf("read algorithmId is: %d", algorithmId)

		hiddenFileSize := uint64(sumSize) - sourceBigFileSize - 8 - 8 - 8

		fp.Seek(int64(hiddenFileSize)*-1-8-8-8, io.SeekEnd)
		log.Warnf("hiddenFileSize: %d", hiddenFileSize)
		outputData := make([]byte, uint64(hiddenFileSize))
		fp.Read(outputData)

		log.Debugf("will write to %s", outputFile)
		binary_utils.WriteBinaryFile(outputFile, outputData)
		log.Debugf("write done.")

	},
}

func init() {
	log.Debugf("init splitCmd")

	rootCmd.AddCommand(splitCmd)

	splitCmd.Flags().StringP("file", "f", "", "filename")
	splitCmd.Flags().StringP("output", "o", "", "filename")
	splitCmd.Flags().StringP("encryption", "e", "none", "aes-128")
	splitCmd.Flags().StringP("password", "p", "", "password")
}
