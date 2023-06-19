package cmd

import (
	"af2f/binary_utils"
	"af2f/common_utils"
	"af2f/encrypt_tool"
	"encoding/binary"
	"github.com/charmbracelet/log"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func validateSplitArgs(file string, output string) bool {
	validated := true
	if fileNotExists(file) {
		log.Errorf("-f file: %s is not exists", file)
		validated = false
	}
	if !fileNotExists(output) {
		log.Errorf("-o file: %s is already exists", output)
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
		password, _ := cmd.Flags().GetString("password")
		if !validateSplitArgs(bigFile, outputFile) {
			log.Errorf("DO NOT PASS THE PARAMS VALIDATE")
			os.Exit(1)
		}

		fp, _ := os.OpenFile(bigFile, os.O_RDONLY, 0644)
		defer fp.Close()

		_, sumSize := binary_utils.ReadBinaryFile(bigFile)

		// read encryptionAlgorithm algorithm
		fp.Seek(-8, io.SeekEnd)
		buffer := make([]byte, 8)
		fp.Read(buffer)
		algorithmId := binary.LittleEndian.Uint64(buffer)
		_, algorithm := common_utils.GetAlgorithmNameById(int(algorithmId))
		log.Debugf("read algorithmId=%d from %s", bigFile, algorithmId)

		// read sourceBigFileSize
		fp.Seek(-16, io.SeekEnd)
		buffer = make([]byte, 8)
		fp.Read(buffer)
		sourceBigFileSize := binary.LittleEndian.Uint64(buffer)
		log.Debugf("read hidden file length = %d from %s", sourceBigFileSize, bigFile)

		hiddenFileSize := uint64(sumSize) - sourceBigFileSize - 8 - 8

		fp.Seek(int64(hiddenFileSize)*-1-8-8, io.SeekEnd)
		log.Warnf("hiddenFileSize: %d", hiddenFileSize)
		outputData := make([]byte, uint64(hiddenFileSize))
		fp.Read(outputData)

		if password != "" {
			log.Debugf("decrypting with %s by %s", algorithm, password)
			outputData = encrypt_tool.AESDecrypt(outputData, algorithm, password)
		} else {
			log.Debug("decrypting without password")
		}

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
	splitCmd.Flags().StringP("password", "p", "", "password")
}
