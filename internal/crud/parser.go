package crud

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/not-for-prod/speedrun/internal/crud/models"
)

// Parse the struct from the source file and return field names and types
func parse(filePath, structName string) ([]models.Field, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var fields []models.Field

	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.Name == structName {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						for _, field := range structType.Fields.List {
							for _, fieldName := range field.Names {
								fields = append(fields, models.Field{
									Name: fieldName.Name,
									Type: fmt.Sprintf("%s", field.Type),
								})
							}
						}
					}
				}
			}
		}
	}

	return fields, nil
}
