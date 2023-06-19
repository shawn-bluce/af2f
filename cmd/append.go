package cmd

import (
	"af2f/binary_utils"
	"af2f/common_utils"
	"af2f/encrypt_tool"
	"encoding/binary"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

func validateAppendArgs(file string, append string, password string, algorithm string) bool {
	validated := true
	if fileNotExists(file) {
		log.Errorf("-f file: %s is not exists", file)
		validated = false
	}
	if fileNotExists(append) {
		log.Errorf("-a file: %s is not exists", append)
		validated = false
	}
	if len(password) > 32 {
		log.Errorf("password lens must less than 32")
		validated = false
	}

	algorithmNameFound := false
	for name, _ := range common_utils.GetAlgorithmMap() {
		if algorithm == name {
			algorithmNameFound = true
			break
		}
	}
	if !algorithmNameFound {
		log.Errorf("algorithm name: %s are not supported", algorithm)
		validated = false
	}

	return validated
}

var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append a file to another file",
	Run: func(cmd *cobra.Command, args []string) {

		bigFile, _ := cmd.Flags().GetString("file")
		appendFile, _ := cmd.Flags().GetString("append")
		encryptionAlgorithm, _ := cmd.Flags().GetString("encryption")
		password, _ := cmd.Flags().GetString("password")
		if !validateAppendArgs(bigFile, appendFile, password, encryptionAlgorithm) {
			log.Errorf("Do not pass the params validate")
			os.Exit(1)
		}

		_, bigfileSize := binary_utils.ReadBinaryFile(bigFile)

		appendData, _ := binary_utils.ReadBinaryFile(appendFile)

		if password != "" {
			appendData = encrypt_tool.AESEncrypt(appendData, encryptionAlgorithm, password)
		}

		log.Infof("append: %s, file-size: %d", appendFile, len(appendData))
		binary_utils.AppendBinaryFile(bigFile, appendData) // append file

		// offset
		bigfileSizeByteArray := make([]byte, 8)
		binary.LittleEndian.PutUint64(bigfileSizeByteArray, uint64(bigfileSize))
		log.Infof("append: bigfileSize, 8bytes, value is %d", bigfileSize)
		binary_utils.AppendBinaryFile(bigFile, bigfileSizeByteArray) // append offset value

		// encryption algorithm
		find, algorithmId := common_utils.GetAlgorithmIdByName(encryptionAlgorithm)
		if !find {
			os.Exit(1)
		}
		algorithmByteArray := make([]byte, 8)
		binary.LittleEndian.PutUint64(algorithmByteArray, uint64(algorithmId))
		log.Infof("append: encryption algorithm, 8bytes")
		binary_utils.AppendBinaryFile(bigFile, algorithmByteArray) // append offset value
	},
}

func init() {
	log.Debugf("init appendCmd")

	rootCmd.AddCommand(appendCmd)

	appendCmd.Flags().StringP("file", "f", "", "filename")
	appendCmd.Flags().StringP("append", "a", "", "filename")
	appendCmd.Flags().StringP("encryption", "e", "aes-128", "aes-128/aes-192/aes-256")
	appendCmd.Flags().StringP("password", "p", "", "password")
}
