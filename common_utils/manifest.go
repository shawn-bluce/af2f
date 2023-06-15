package common_utils

import (
	"encoding/json"
	"os"
)

type Manifest struct {
	Version int `json:"version"`
}

func GetManifest() Manifest {
	file, _ := os.Open("manifest.json")
	defer file.Close()

	var manifest Manifest

	decoder := json.NewDecoder(file)
	decoder.Decode(&manifest)

	return manifest
}
