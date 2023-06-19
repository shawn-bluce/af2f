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
	for name, id := range GetAlgorithmMap() {
		if algorithmName == name {
			log.Debugf("find the algorithm id %d by name %s", id, name)
			return true, id
		}
	}
	return false, -1
}

func GetAlgorithmNameById(algorithmId int) (bool, string) {
	for name, id := range GetAlgorithmMap() {
		if algorithmId == id {
			log.Infof("find the algorithm %s by id %d", name, algorithmId)
			return true, name
		}
	}
	return false, "NOT_FOUND"
}
