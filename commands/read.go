package commands

import (
	"fmt"
	"strings"
)

type ReadCommand struct {
	Meta
}

func (c *ReadCommand) Help() string {
	helpText := `
Usage: aero read <key>
  Returns the value of the metdata key provided.
  `

	return strings.TrimSpace(helpText)
}

func (c *ReadCommand) Synopsis() string {
	return "Read a metadata value"
}

func (c *ReadCommand) Run(args []string) int {
	if len(args) < 1 {
		fmt.Printf("Expected a metdata key to read.\n")
		return 1
	}

	metadataKey := args[0]

	p := *c.Meta.CurrentProvider
	metadataValue, err := p.Read(metadataKey)

	if err != nil {
		fmt.Printf("error reading value [%s]: %s \n", metadataKey, err)
		return 1
	}

	if metadataValue == "" {
		fmt.Printf("No known metadata [%s].\n", metadataKey)
		return 1
	}

	fmt.Printf("%s\n", metadataValue)
	return 0
}
