package model

import (
	"fmt"
	"path/filepath"

	"github.com/not-for-prod/speedrun/cmd/crud/generators"
	"github.com/not-for-prod/speedrun/cmd/crud/models"
	"github.com/not-for-prod/speedrun/internal/pkg/string-tools"
)

var modelTemplate = `
package {{toSnakeCase .StructName}}_repository

import (
	{{ range $import := .Imports -}}
		{{ $import.Alias }} "{{ $import.Pkg }}"
	{{ end }}
)

type db{{.StructName}} struct {
{{range $index, $field := .Fields}}    {{$field.Name}} {{$field.Type}} ` + "`db:\"{{toSnakeCase $field.Name}}\"`" + `
{{end}}}

func fromEntity(e {{toSnakeCase .StructName}}.{{.StructName}}) db{{.StructName}} {
	return db{{.StructName}}{
	{{range $index, $field := .Fields}}    {{$field.Name}}: e.{{$field.Name}},
	{{end}}}
}

func (m db{{.StructName}}) toEntity() {{toSnakeCase $.StructName}}.{{.StructName}} {
	return {{toSnakeCase .StructName}}.{{.StructName}}{
	{{range $index, $field := .Fields}}    {{$field.Name}}: m.{{$field.Name}},
	{{end}}}
}
`

type generateData struct {
	StructName string
	Imports    []models.Import
	Fields     []models.Field
}

func Generate(packageMod, src, dst, structName string, fields []models.Field) []models.File {
	gData := generateData{
		StructName: structName,
		Imports: []models.Import{
			{Alias: string_tools.SnakeCase(structName), Pkg: packageMod + "/" + filepath.Dir(src)},
		},
		Fields: fields,
	}

	generated := generators.Generate(
		fmt.Sprintf("%s/%s/models.go", dst, string_tools.SnakeCase(structName)),
		modelTemplate,
		gData,
	)

	return []models.File{generated}
}
