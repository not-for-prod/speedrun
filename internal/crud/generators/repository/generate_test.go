package repository

import (
	"testing"

	"github.com/not-for-prod/speedrun/internal/pkg/logger"
)

func TestGenerate(t *testing.T) {
	data := data{
		StructName: "aboba",
	}

	result := Generate(data)
	logger.Info(result)
}
