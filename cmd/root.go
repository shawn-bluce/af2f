package cmd

import (
	"github.com/charmbracelet/log"
	"os"

	"github.com/spf13/cobra"
)

func fileNotExists(filename string) bool {
	log.Debugf("check file: %s exists", filename)
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Debugf("file: %s is not exists", filename)
		return true
	}
	isDir := info.IsDir()
	log.Debugf("file: %s is exists, but it's dir", filename)
	return isDir
}

var rootCmd = &cobra.Command{
	Use:   "af2f",
	Short: "Hidden a file to another file.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.Debugf("init rootCmd")
}
