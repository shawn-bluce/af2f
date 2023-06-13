/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"af2f/cmd"
	"github.com/charmbracelet/log"
	"os"
)

func main() {
	logLevel := os.Getenv("AF2F_LOG_LEVEL")
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	log.Debugf("Pre cmd.Execute()")
	cmd.Execute()
	log.Debugf("Post cmd.Execute()")
}
