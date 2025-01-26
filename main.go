package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// Helper function to convert camelCase to snake_case
func toSnakeCase(str string) string {
	var result []rune
	for i, c := range str {
		if i > 0 && 'A' <= c && c <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, rune(strings.ToLower(string(c))[0]))
	}
	return string(result)
}

// Parse the struct from the source file and return field names and types
func parseStruct(filePath, structName string) ([]string, []string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	var fieldNames []string
	var fieldTypes []string

	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok && typeSpec.Name.Name == structName {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						for _, field := range structType.Fields.List {
							for _, fieldName := range field.Names {
								fieldNames = append(fieldNames, fieldName.Name)

								// Get the field type (we assume it's simple types like string, int, etc.)
								fieldTypes = append(fieldTypes, fmt.Sprintf("%s", field.Type))
							}
						}
					}
				}
			}
		}
	}

	return fieldNames, fieldTypes, nil
}

// Create the directory if it doesn't exist
func createDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// Write a file from template
func writeFileFromTemplate(filename string, tmpl *template.Template, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

// Struct representing the template data for models.go
type ModelData struct {
	StructName string
	Fields     []string
	FieldTypes []string
}

// Struct representing the template data for repository functions
type RepoData struct {
	StructName string
	Fields     []string
	FieldTypes []string
}

// Define the Cobra command
func main() {
	var src string
	var dst string

	// Define the Cobra CLI command
	var rootCmd = &cobra.Command{
		Use:   "crud",
		Short: "Generates CRUD code for a Go struct",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse the input src argument
			parts := strings.Split(src, "::")
			if len(parts) != 3 {
				log.Fatalf("Invalid --src argument, expected format: <file>::<struct>::<identity>")
			}

			filePath, structName, _ := parts[0], parts[1], parts[2]
			fmt.Printf("Parsing struct '%s' from file '%s'...\n", structName, filePath)

			// Parse the struct to get the fields and types
			fieldNames, fieldTypes, err := parseStruct(filePath, structName)
			if err != nil {
				log.Fatalf("Error parsing struct: %v", err)
			}

			// Create the destination folder structure
			structFolder := strings.ToLower(strings.ReplaceAll(structName, "", "-"))
			dstFolder := filepath.Join(dst, structFolder)

			err = createDirIfNotExists(dstFolder)
			if err != nil {
				log.Fatalf("Failed to create destination folder: %v", err)
			}

			// Create the models.go file
			modelData := ModelData{
				StructName: structName,
				Fields:     fieldNames,
				FieldTypes: fieldTypes,
			}

			modelFile, err := os.Create(filepath.Join(dstFolder, "models.go"))
			if err != nil {
				log.Fatalf("Error creating models.go file: %v", err)
			}
			defer modelFile.Close()

			modelTemplate := `
package {{.StructName}}

type {{.StructName}} struct {
{{range $index, $field := .Fields}}    {{$field}} {{$fieldType}} \`json:"{{$field}}" db:"{{toSnakeCase $field}}"\`
{{end}}
}
`
			// Parse and execute the template for models.go
			tmpl, err := template.New("models").Funcs(template.FuncMap{"toSnakeCase": toSnakeCase}).Parse(modelTemplate)
			if err != nil {
				log.Fatalf("Error parsing template: %v", err)
			}

			err = tmpl.Execute(modelFile, modelData)
			if err != nil {
				log.Fatalf("Error executing template: %v", err)
			}

			// Generate repository.go file
			repoData := RepoData{
				StructName: structName,
				Fields:     fieldNames,
				FieldTypes: fieldTypes,
			}

			repoFile, err := os.Create(filepath.Join(dstFolder, "repository.go"))
			if err != nil {
				log.Fatalf("Error creating repository.go file: %v", err)
			}
			defer repoFile.Close()

			repoTemplate := `
package {{.StructName}}

import (
    "fmt"
    // Add other imports here (e.g., database/sql)
)

func Get{{.StructName}}ByID(id string) (*{{.StructName}}, error) {
    // Example SQL query logic depending on ID field type (e.g., int, string, etc.)
    return nil, fmt.Errorf("not implemented")
}
`

			// Parse and execute the template for repository.go
			repoTmpl, err := template.New("repository").Parse(repoTemplate)
			if err != nil {
				log.Fatalf("Error parsing template: %v", err)
			}

			err = repoTmpl.Execute(repoFile, repoData)
			if err != nil {
				log.Fatalf("Error executing repository template: %v", err)
			}

			// Generate other CRUD operation files (create.go, get.go, update.go, delete.go)
			crudFiles := []struct {
				filename string
				template string
			}{
				{"create.go", "func Create{{.StructName}}(data *{{.StructName}}) error { return nil }"},
				{"get.go", "func Get{{.StructName}}(id string) (*{{.StructName}}, error) { return nil, nil }"},
				{"update.go", "func Update{{.StructName}}(data *{{.StructName}}) error { return nil }"},
				{"delete.go", "func Delete{{.StructName}}(id string) error { return nil }"},
			}

			for _, crudFile := range crudFiles {
				file, err := os.Create(filepath.Join(dstFolder, crudFile.filename))
				if err != nil {
					log.Fatalf("Error creating %s file: %v", crudFile.filename, err)
				}
				defer file.Close()

				crudTemplate := fmt.Sprintf(`
package {{.StructName}}

import "fmt"

{{.template}}
`, crudFile.template)

				crudTmpl, err := template.New(crudFile.filename).Parse(crudTemplate)
				if err != nil {
					log.Fatalf("Error parsing %s template: %v", crudFile.filename, err)
				}

				err = crudTmpl.Execute(file, repoData)
				if err != nil {
					log.Fatalf("Error executing %s template: %v", crudFile.filename, err)
				}
			}

			fmt.Printf("CRUD code generated successfully in folder '%s'\n", dstFolder)
		},
	}

	// Define flags for CLI arguments
	rootCmd.PersistentFlags().StringVar(&src, "src", "", "Source file with struct definition (required)")
	rootCmd.PersistentFlags().StringVar(&dst, "dst", ".", "Destination directory for generated files (default is current directory)")
	rootCmd.MarkPersistentFlagRequired("src")

	// Execute the command
	if err := rootCmd.Execute(); err
