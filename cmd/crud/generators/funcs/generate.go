package funcs

import (
	"fmt"
	"path/filepath"

	"github.com/not-for-prod/speedrun/cmd/crud/generators"
	"github.com/not-for-prod/speedrun/cmd/crud/generators/funcs/templates"
	"github.com/not-for-prod/speedrun/cmd/crud/models"
	string_tools "github.com/not-for-prod/speedrun/internal/pkg/string-tools"
	"github.com/samber/lo"
)

var sql = `
package sql

import (
	_ "embed"
)

//go:embed create.sql
var Create string

//go:embed get.sql
var Get string

//go:embed update.sql
var Update string

//go:embed delete.sql
var Delete string
`

type data struct {
	StructName string
	ID         models.Field
	Imports    []models.Import
	Fields     []models.Field
}

func Generate(packageMod, src, dst, structName string, id models.Field, fields []models.Field) []models.File {
	data := data{
		StructName: structName,
		ID:         id,
		Fields:     append([]models.Field{id}, fields...),
		Imports: []models.Import{
			{Alias: string_tools.SnakeCase(structName), Pkg: packageMod + "/" + filepath.Dir(src)},
			{Alias: "sql", Pkg: packageMod + "/" + dst + "/" + string_tools.SnakeCase(structName) + "/sql"},
		},
	}

	data.Fields = lo.Filter(data.Fields, func(f models.Field, _ int) bool {
		return f.Name != data.ID.Name
	})

	data.Fields = append([]models.Field{data.ID}, data.Fields...)

	files := make([]models.File, 0, 4*2+1)
	generated := generators.Generate(
		fmt.Sprintf("%s/%s/sql/sql.go", dst, string_tools.SnakeCase(structName)),
		sql,
		data,
	)
	files = append(files, generated)

	// C
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/create.go", dst, string_tools.SnakeCase(structName)),
		templates.CreateFuncTemplate,
		data,
	)
	files = append(files, generated)
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/sql/create.sql", dst, string_tools.SnakeCase(structName)),
		templates.CreateSQL,
		data,
	)
	files = append(files, generated)

	// R
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/get.go", dst, string_tools.SnakeCase(structName)),
		templates.GetFuncTemplate,
		data,
	)
	files = append(files, generated)
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/sql/get.sql", dst, string_tools.SnakeCase(structName)),
		templates.GetSQL,
		data,
	)
	files = append(files, generated)

	// U
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/update.go", dst, string_tools.SnakeCase(structName)),
		templates.UpdateFuncTemplate,
		data,
	)
	files = append(files, generated)
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/sql/update.sql", dst, string_tools.SnakeCase(structName)),
		templates.UpdateSQL,
		data,
	)
	files = append(files, generated)

	// D
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/delete.go", dst, string_tools.SnakeCase(structName)),
		templates.DeleteFuncTemplate,
		data,
	)
	files = append(files, generated)
	generated = generators.Generate(
		fmt.Sprintf("%s/%s/sql/delete.sql", dst, string_tools.SnakeCase(structName)),
		templates.DeleteSQL,
		data,
	)
	files = append(files, generated)

	return files
}
