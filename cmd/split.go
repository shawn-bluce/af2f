package cmd

import (
	"github.com/charmbracelet/log"
	"os"

	"github.com/spf13/cobra"
)

func validateSplitArgs(file string, output string, password string) bool {
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

	return validated
}

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a file from a file build from af2f",
	Run: func(cmd *cobra.Command, args []string) {
		bigFile, _ := cmd.Flags().GetString("file")
		outputFile, _ := cmd.Flags().GetString("output")
		password, _ := cmd.Flags().GetString("password")
		if !validateSplitArgs(bigFile, outputFile, password) {
			log.Errorf("Do not pass the params validate")
			os.Exit(1)
		}

	},
}

func init() {
	log.Debugf("init splitCmd")

	rootCmd.AddCommand(splitCmd)

	splitCmd.Flags().StringP("file", "f", "", "filename")
	splitCmd.Flags().StringP("output", "o", "", "filename")
	splitCmd.Flags().StringP("password", "p", "", "password")
}
