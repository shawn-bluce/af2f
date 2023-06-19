package common_utils

import (
	"github.com/charmbracelet/log"
)

func GetAlgorithmMap() map[string]int {
	return map[string]int{
		"none":    0,
		"aes-128": 1,
		"aes-192": 2,
		"aes-256": 3,
	}
}

func GetAlgorithmIdByName(algorithmName string) (bool, int) {
	find := false
	for name, id := range GetAlgorithmMap() {
		if algorithmName == name {
			find = true
			log.Infof("find the algorithm id %d by name %s", id, name)
			break
		}
	}
	if find {
		return true, GetAlgorithmMap()["aes-128"]
	} else {
		log.Errorf("Not found algorithm name: %s", algorithmName)
		return false, -1
	}
}

func GetAlgorithmNameById(algorithmId int) (bool, string) {
	for name, id := range GetAlgorithmMap() {
		if algorithmId == id {
			log.Infof("find the algorithm %s by id %d", name, algorithmId)
			return true, name
		}
	}
	log.Errorf("Not found algorithm id: %d", algorithmId)
	return false, "NOT_FOUND"
}
