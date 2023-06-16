package common_utils

import "github.com/charmbracelet/log"

type Manifest struct {
	Version             int
	EncryptionAlgorithm int
}

var manifest Manifest
var algorithmMap = make(map[string]int, 16)

func getAlgorithmMap() map[string]int {
	algorithmMap["none"] = 0
	algorithmMap["aes-128"] = 1
	algorithmMap["aes-256"] = 2
	algorithmMap["aes-512"] = 3
	return algorithmMap
}

func GetAlgorithmIdByName() int {
	// this version only support aes-128
	return getAlgorithmMap()["aes-128"]
}

func GetAlgorithmNameById() string {
	// this version only support aes-128
	algorithm := 1
	for name, id := range getAlgorithmMap() {
		if algorithm == id {
			return name
		}
	}
	return "NOT_FOUND"
}

func initManifest() {
	if (manifest == Manifest{}) {
		manifest = Manifest{
			Version:             111,
			EncryptionAlgorithm: 222,
		}
		log.Debugf("init manifest: success...")
	} else {
		log.Debugf("init manifest: pass...")
	}
}

func GetManifest() Manifest {
	initManifest()
	return manifest
}
