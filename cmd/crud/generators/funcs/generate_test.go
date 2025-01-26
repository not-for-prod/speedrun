package funcs

import (
	"testing"

	"github.com/not-for-prod/speedrun/cmd/crud/models"
	"github.com/not-for-prod/speedrun/internal/pkg/logger"
)

func TestGenerate(t *testing.T) {
	result := Generate(
		"github.com/not-for-prod/speedrun/cmd/crud/example/in/peach.go",
		"github.com/not-for-prod/speedrun/cmd/crud/example/out/infrastructure",
		"Aboba",
		models.Field{Name: "ID", Type: "int"},
		[]models.Field{{Name: "Name", Type: "string"}, {Name: "Age", Type: "int"}},
	)
	logger.Info(result)
}
