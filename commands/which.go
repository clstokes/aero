package commands

import (
	"fmt"
	"strings"
)

type WhichCommand struct {
	Meta
}

func (c *WhichCommand) Help() string {
	helpText := `
Usage: aero which
  Returns the name of the provider of the current instance.
  `

	return strings.TrimSpace(helpText)
}

func (c *WhichCommand) Synopsis() string {
	return "Determine the current provider"
}

func (c *WhichCommand) Run(args []string) int {
	p := *c.Meta.CurrentProvider
	fmt.Printf("%s", p.Name())
	return 0
}
