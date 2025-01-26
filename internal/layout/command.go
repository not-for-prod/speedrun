package layout

import "os"

type layoutCommand struct {
	serviceName string
}

func (c layoutCommand) execute() {
	layout := []string{
		"cmd/" + c.serviceName,
		"internal/config",
		"internal/" + c.serviceName + "/app",
		"internal/" + c.serviceName + "/domain/entity",
		"internal/" + c.serviceName + "/domain/valueobject",
		"internal/" + c.serviceName + "/domain/microtype",
		"internal/" + c.serviceName + "/infrastructure",
		"internal/pkg",
		"pkg",
	}

	for _, path := range layout {
		os.MkdirAll(path, os.ModePerm)
	}
}
