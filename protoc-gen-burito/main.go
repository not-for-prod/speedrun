package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if file.Generate {
				generateFile(file, plugin)
			}
		}

		return nil
	})
}

func generateFile(f *protogen.File, p *protogen.Plugin) {
	g := p.NewGeneratedFile(f.GeneratedFilenamePrefix+".pb.burito.go", f.GoImportPath)
	g.P("package ", f.GoPackageName)
	g.P()
}
