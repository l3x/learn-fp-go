package infrastructure_test

import (
	"04_onion/infrastructure"
	"testing"
)

func TestLocalInteractor(t *testing.T) {
	localInteractor, err := infrastructure.GetLocalInteractor()
	if err != nil {
		t.Error("GetLocalInteractor failed")
	}
	localInteractor.LocalRepository.FileExists()
}

func TestFileExists(t *testing.T) {
	t.Error("We haven't written our test yet")
}
