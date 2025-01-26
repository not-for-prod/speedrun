package templates

var CreateSQL = `insert into {{toSnakeCase .StructName}} 
({{range $index, $field := .Fields}}{{if $index}}, {{end}}{{toSnakeCase $field.Name}}{{end}})
values ({{range $i, $field := .Fields}}{{if $i}}, {{end}}${{$i | inc}}{{end}})
returning {{toSnakeCase .ID.Name}};
`

var GetSQL = `select ({{range $index, $field := .Fields}}{{if $index}}, {{end}}{{toSnakeCase $field.Name}}{{end}})
from {{toSnakeCase .StructName}} 
where {{toSnakeCase .ID.Name}} = $1;
`

var UpdateSQL = `update {{toSnakeCase .StructName}}
set {{range $index, $field := .Fields}}{{if $index}}, {{end}}
    {{toSnakeCase $field.Name}} = ${{$index | inc}}{{end}}
where {{toSnakeCase .ID.Name}} = $1;
`

var DeleteSQL = `delete from {{toSnakeCase .StructName}}
where {{toSnakeCase .ID.Name}} = $1;
`
