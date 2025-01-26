package migrations

import (
	"testing"

	"github.com/not-for-prod/speedrun/internal/crud/models"
	"github.com/not-for-prod/speedrun/internal/pkg/logger"
)

func TestGenerate(t *testing.T) {
	result := Generate(
		"aboba",
		models.Field{Name: "Id", Type: "int"},
		[]models.Field{
			{Name: "Name", Type: "string"},
			{Name: "Age", Type: "float64"},
		},
	)
	logger.Info(result)
}
