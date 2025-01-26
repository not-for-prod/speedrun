package crud

import (
	"strings"

	"github.com/not-for-prod/speedrun/internal/crud/generators"
	"github.com/not-for-prod/speedrun/internal/crud/generators/funcs"
	"github.com/not-for-prod/speedrun/internal/crud/generators/model"
	"github.com/not-for-prod/speedrun/internal/crud/generators/repository"
	"github.com/not-for-prod/speedrun/internal/crud/models"
	"github.com/not-for-prod/speedrun/internal/pkg/logger"
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

	for _, f := range fields {
		if f.Name == c.idField {
			id = f
			break
		}
	}

	logger.Info("id:", id)

	// generate models file
	files := make([]models.File, 0, 10)
	files = append(
		model.Generate(currentPackage, c.srcPath, c.dstPath, c.structName, fields),
		append(repository.Generate(c.dstPath, c.structName),
			funcs.Generate(currentPackage, c.srcPath, c.dstPath, c.structName, id, fields)...,
		)...,
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
