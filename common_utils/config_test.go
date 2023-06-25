package common_utils

import "testing"

func TestGetAlgorithmIdByName(t *testing.T) {
	algorithmList := []string{"none", "aes-128", "aes-192", "aes-256"}

	for id, algorithm := range algorithmList {
		find, resultId := GetAlgorithmIdByName(algorithm)
		if !find || id != resultId {
			t.Errorf("get %s ERROR", algorithm)
		}
	}

	find, resultId := GetAlgorithmIdByName("aes-233")
	if find || resultId >= 0 {
		t.Errorf("aes-233 should be not found")
	}

}

func TestGetAlgorithmNameById(t *testing.T) {
	algorithmList := []string{"none", "aes-128", "aes-192", "aes-256"}

	for id, algorithm := range algorithmList {
		find, resultAlgorithm := GetAlgorithmNameById(id)
		if !find || algorithm != resultAlgorithm {
			t.Errorf("get id=%d ERROR", id)
		}
	}

	find, resultAlgorithm := GetAlgorithmNameById(233)
	if find || resultAlgorithm != "NOT_FOUND" {
		t.Errorf("id 233 should be not found")
	}
}
