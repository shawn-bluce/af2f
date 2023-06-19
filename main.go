/*
Copyright © 2023 Shawn <shawnbluce@gmail.com>
*/
package main

import (
	"af2f/cmd"
	"github.com/charmbracelet/log"
	"os"
)

func main() {
	debug := os.Getenv("AF2F_DEBUG") == "1"
	if debug {
		log.SetLevel(log.DebugLevel)
		log.Infof("OS ENV AF2F_DEBUG=1, set logLevel=debug")
	}
	log.Debugf("Pre cmd.Execute()")
	cmd.Execute()
	log.Debugf("Post cmd.Execute()")
}
