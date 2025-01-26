package generators

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/not-for-prod/speedrun/internal/crud/models"
	string_tools "github.com/not-for-prod/speedrun/internal/pkg/string-tools"
)

func inc(i int) int {
	return i + 1
}

func Generate(name, modelTemplate string, data any) models.File {
	// Parse and execute the template for models.go
	tmpl, err := template.New(name).
		Funcs(template.FuncMap{
			"toSnakeCase": string_tools.SnakeCase,
			"inc":         inc,
		}).Parse(modelTemplate)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

	return models.File{
		Path: name,
		Data: buf.String(),
	}
}

// WriteToFile writes r to the file with path
func WriteToFile(path string, r io.Reader) error {
	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, r)
	return err
}

func WriteStringToFile(path string, s string) error {
	return WriteToFile(path, strings.NewReader(s))
}
