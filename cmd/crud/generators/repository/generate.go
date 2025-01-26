package repository

import (
	"fmt"

	"github.com/not-for-prod/speedrun/cmd/crud/generators"
	"github.com/not-for-prod/speedrun/cmd/crud/models"
	string_tools "github.com/not-for-prod/speedrun/internal/pkg/string-tools"
)

var repositoryTemplate = `package {{toSnakeCase .StructName}}_repository

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
)

type {{.StructName}}Repository struct {
	db        *sqlx.DB
	ctxGetter *trmsqlx.CtxGetter
	txManager *manager.Manager
}

func New(
	db *sqlx.DB,
	ctxGetter *trmsqlx.CtxGetter,
	txManager *manager.Manager,
) *{{.StructName}}Repository {
	return &{{.StructName}}Repository{
		db:        db,
		ctxGetter: ctxGetter,
		txManager: txManager,
	}
}
`

type data struct {
	StructName string
}

func Generate(dst, structName string) []models.File {
	generated := generators.Generate(
		fmt.Sprintf("%s/%s/repository.go", dst, string_tools.SnakeCase(structName)),
		repositoryTemplate,
		data{StructName: structName},
	)

	return []models.File{generated}
}
