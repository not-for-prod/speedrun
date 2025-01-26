package crud

import (
	"io/ioutil"

	"github.com/not-for-prod/speedrun/internal/pkg/logger"
	"golang.org/x/mod/modfile"
)

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		logger.Fatalf("failed to read go.mod:", err.Error())
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
}
