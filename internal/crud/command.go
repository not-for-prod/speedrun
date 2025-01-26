package crud

import (
	"strings"

	"github.com/not-for-prod/speedrun/internal/crud/generators"
	"github.com/not-for-prod/speedrun/internal/crud/generators/funcs"
	"github.com/not-for-prod/speedrun/internal/crud/generators/migrations"
	"github.com/not-for-prod/speedrun/internal/crud/generators/model"
	"github.com/not-for-prod/speedrun/internal/crud/generators/repository"
	"github.com/not-for-prod/speedrun/internal/crud/models"
	"github.com/not-for-prod/speedrun/internal/pkg/logger"
	"github.com/samber/lo"
)

type generationCommand struct {
	srcPath    string
	dstPath    string
	structName string
	idField    string
}

func generationCommandFromString(src, dst string) generationCommand {
	if src == "" {
		logger.Fatalf("src flag is required")
	}

	if dst == "" {
		logger.Fatalf("dst flag is required")
	}

	parts := strings.Split(src, "::")
	if len(parts) != 3 {
		logger.Fatalf("invalid path format expected format: <path to file>::<struct name>::<id field>")
	}

	return generationCommand{
		srcPath:    parts[0],
		dstPath:    dst,
		structName: parts[1],
		idField:    parts[2],
	}
}

func (c generationCommand) execute() {
	currentPackage := GetModuleName()
	logger.Info("current package:", currentPackage)

	fields, err := parse(c.srcPath, c.structName)
	if err != nil {
		logger.Fatalf("failed to parse struct: %v", err)
	}

	logger.Info("src:", c.srcPath, "dst:", c.dstPath)
	logger.Info("fields:", fields)

	var id models.Field

	fields = lo.Filter(fields, func(item models.Field, _ int) bool {
		if item.Name == c.idField {
			id = item
			return false
		}

		return true
	})

	logger.Info("id:", id)

	// generate models file
	files := make([]models.File, 0, 4*2+2+1)
	files = append(
		append(
			model.Generate(currentPackage, c.srcPath, c.dstPath, c.structName, id, fields),
			append(repository.Generate(c.dstPath, c.structName),
				funcs.Generate(currentPackage, c.srcPath, c.dstPath, c.structName, id, fields)...,
			)...,
		), migrations.Generate(c.structName, id, fields)...,
	)

	logger.Info("generated:", len(files), "files")

	for _, file := range files {
		err = generators.WriteStringToFile(file.Path, file.Data)
		if err != nil {
			logger.Fatalf("failed to write file: %v", err)
		}

		logger.Info("wrote:", file.Path)
	}
}
