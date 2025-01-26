package models

type Field struct {
	Name string
	Type string
}

type Import struct {
	Alias string
	Pkg   string
}

type File struct {
	Path string
	Data string
}
