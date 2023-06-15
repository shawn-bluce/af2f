package cmd

import (
	"af2f/binary_utils"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

func validateAppendArgs(file string, append string, password string) bool {
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

	return validated
}

var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append a file to another file",
	Run: func(cmd *cobra.Command, args []string) {

		bigFile, _ := cmd.Flags().GetString("file")
		appendFile, _ := cmd.Flags().GetString("append")
		password, _ := cmd.Flags().GetString("password")
		if !validateAppendArgs(bigFile, appendFile, password) {
			log.Errorf("Do not pass the params validate")
			os.Exit(1)
		}

		data := binary_utils.ReadBinaryFile(appendFile)
		binary_utils.AppendBinaryFile(bigFile, data)
	},
}

func init() {
	log.Debugf("init appendCmd")

	rootCmd.AddCommand(appendCmd)

	appendCmd.Flags().StringP("file", "f", "", "filename")
	appendCmd.Flags().StringP("append", "a", "", "filename")
	appendCmd.Flags().StringP("password", "p", "", "password")
}
