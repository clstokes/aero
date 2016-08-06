package commands

import (
	"fmt"
	"github.com/clstokes/aero/structs"
	"sort"
	"strings"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Help() string {
	helpText := `
Usage: aero list
  List available metadata keys. May include keys that are not supported
  on all providers.
  `

	return strings.TrimSpace(helpText)
}

func (c *ListCommand) Synopsis() string {
	return "List available metadata keys"
}

func (c *ListCommand) Run(args []string) int {
	sort.Strings(structs.AllKeys)

	for _, key := range structs.AllKeys {
		fmt.Printf("%s\n", key)
	}

	return 0
}
