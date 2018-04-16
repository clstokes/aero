// +build go1.7

package main

import (
	"fmt"
	"github.com/clstokes/aero/commands"
	"github.com/clstokes/aero/providers"
	"github.com/clstokes/aero/structs"
	"github.com/mitchellh/cli"
	"os"
)

const VERSION = "v0.1.0"

type Factory func(defaults structs.ProviderMapping) structs.Provider

var fingerprinters = map[string]Factory{
	structs.NAME_AMAZON: providers.InitAmazon,
	structs.NAME_GOOGLE: providers.InitGoogle,
}

var currentProvider structs.Provider

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	cli := &cli.CLI{
		Args:     args,
		Commands: commands.Commands(&currentProvider),
		Version:  VERSION,
	}

	// TODO: This is passed above but set here - feels wonky.
	currentProvider = getCurrentProvider()
	if currentProvider == nil {
		fmt.Printf("Unknown provider.\n")
		return 1
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}

func getCurrentProvider() structs.Provider {
	for _, val := range fingerprinters {
		emptyDefaults := new(structs.ProviderMapping)
		provider := val(*emptyDefaults)
		if provider.IsCurrentProvider() {
			return provider
		}
	}
	return nil
}
