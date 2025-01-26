package templates

// C
var CreateFuncTemplate = `
package {{toSnakeCase .StructName}}_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	{{ range $import := .Imports }}
		{{ $import.Alias }} "{{ $import.Pkg }}"
	{{ end }}
)

func (r *{{.StructName}}Repository) Create(ctx context.Context, {{toSnakeCase .StructName}} {{toSnakeCase .StructName}}.{{.StructName}}) ({{.ID.Type}}, error) {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/{{toSnakeCase .StructName}}/create.go")
	defer span.End()

	var dbID {{.ID.Type}}

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx, 
		&dbID,
		sql.Insert,
		{{range $index, $field := .Fields}}{{toSnakeCase $.StructName}}.{{$field.Name}},
		{{end}}
	)	
	if err != nil {

	}

	return dbID, nil
}
`

// R
var GetFuncTemplate = `
package {{toSnakeCase .StructName}}_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	{{ range $import := .Imports -}}
		{{ $import.Alias }} "{{ $import.Pkg }}"
	{{ end }}
)

func (r *{{.StructName}}Repository) Get(ctx context.Context, id {{.ID.Type}}) (.{{.StructName}}, error) {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/{{toSnakeCase .StructName}}/get.go")
	defer span.End()

	var db{{.StructName}} db{{.StructName}}

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx, 
		&db{{.StructName}},
		sql.Get,
		id,
	)	
	if err != nil {

	}

	return db{{.StructName}}.toEntity(), nil
}
`

// U
var UpdateFuncTemplate = `
package {{toSnakeCase .StructName}}_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	{{ range $import := .Imports -}}
		{{ $import.Alias }} "{{ $import.Pkg }}"
	{{ end }}
)

func (r *{{.StructName}}Repository) Update(ctx context.Context, {{toSnakeCase .StructName}} {{toSnakeCase .StructName}}.{{.StructName}}) error {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/{{toSnakeCase .StructName}}/update.go")
	defer span.End()

	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx, 
		sql.Update,
		{{range $index, $field := .Fields}}{{toSnakeCase $.StructName}}.{{$field.Name}},
		{{end}}
	)	
	if err != nil {

	}

	return nil
}
`

// D
var DeleteFuncTemplate = `
package {{toSnakeCase .StructName}}_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	{{ range $import := .Imports -}}
		{{ $import.Alias }} "{{ $import.Pkg }}"
	{{ end }}
)

func (r *{{.StructName}}Repository) Delete(ctx context.Context, id {{.ID.Type}}) error {
	ctx, span := otel.Tracer("").Start(ctx, "/repository/{{toSnakeCase .StructName}}/get.go")
	defer span.End()

	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		sql.Delete,
		id,
	)	
	if err != nil {
		
	}

	return nil
}
`
