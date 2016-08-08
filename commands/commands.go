package commands

import (
	"github.com/clstokes/aero/structs"
	"github.com/mitchellh/cli"
)

type Meta struct {
	CurrentProvider *structs.Provider
}

func Commands(currentProvider *structs.Provider) map[string]cli.CommandFactory {
	meta := Meta{
		CurrentProvider: currentProvider,
	}

	commands := map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &ListCommand{
				Meta: meta,
			}, nil
		},
		"read": func() (cli.Command, error) {
			return &ReadCommand{
				Meta: meta,
			}, nil
		},
	}

	return commands
}
