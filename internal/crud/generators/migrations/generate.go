package migrations

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/not-for-prod/speedrun/internal/crud/generators"
	"github.com/not-for-prod/speedrun/internal/crud/models"
	string_tools "github.com/not-for-prod/speedrun/internal/pkg/string-tools"
)

var (
	migrationUp = `
create table {{toSnakeCase .StructName}} (
(
    {{toSnakeCase .ID.Name}} {{sqlTypeMap .ID.Type}} PRIMARY KEY,
    {{range $i, $field := .Fields}}{{toSnakeCase $field.Name}} {{sqlTypeMap $field.Type}}{{if $i}}, {{end}}
	{{end}}
);
`
	migrationDown = `
drop table if exists {{toSnakeCase .StructName}};
`
)

type data struct {
	StructName string
	ID         models.Field
	Fields     []models.Field
}

func Generate(structName string, id models.Field, fields []models.Field) []models.File {
	up := generators.Generate(
		fmt.Sprintf("migrations/%s_%s_table.up.sql", uuid.NewString(), string_tools.SnakeCase(structName)),
		migrationUp,
		data{
			StructName: structName,
			ID:         id,
			Fields:     fields,
		},
	)

	down := generators.Generate(fmt.Sprintf("migrations/%s_%s_table.down.sql", uuid.NewString(), string_tools.SnakeCase(structName)),
		migrationDown,
		data{
			StructName: structName,
			ID:         id,
			Fields:     fields,
		},
	)

	return []models.File{up, down}
}
