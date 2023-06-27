package cmd

import (
	"encoding/binary"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func validateRestoreArgs(file string) bool {
	if fileNotExists(file) {
		log.Errorf("-f file: %s is not exists", file)
		return false
	}
	return true
}

func doRestore(filename string) {

	if !validateRestoreArgs(filename) {
		log.Errorf("DO NOT PASS THE PARAMS VALIDATE")
		os.Exit(1)
	}

	info, _ := os.Stat(filename)

	fp, _ := os.OpenFile(filename, os.O_RDWR, 0644)
	defer fp.Close()

	fp.Seek(-16, io.SeekEnd)
	buffer := make([]byte, 8)
	fp.Read(buffer)
	sourceBigFileSize := binary.LittleEndian.Uint64(buffer)

	if sourceBigFileSize < 0 || int64(sourceBigFileSize) > info.Size() {
		log.Errorf("restore fail. maybe this file is not generate by af2f.")
	}

	fp.Truncate(int64(sourceBigFileSize))
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore file(delete appended data)",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		doRestore(filename)
	},
}

func init() {
	log.Debugf("init restoreCmd")

	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().StringP("file", "f", "", "filename")
}
