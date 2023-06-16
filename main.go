/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"af2f/cmd"
	"af2f/common_utils"
	"github.com/charmbracelet/log"
	"os"
)

func main() {
	debug := os.Getenv("AF2F_DEBUG") == "1"
	if debug {
		log.SetLevel(log.DebugLevel)
		log.Infof("OS ENV AF2F_DEBUG=1, set logLevel=debug")
		log.Infof("Now Version is %d", common_utils.GetManifest().Version)
	}
	log.Debugf("Pre cmd.Execute()")
	cmd.Execute()
	log.Debugf("Post cmd.Execute()")
}
