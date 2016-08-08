package main

import (
	"fmt"
	"github.com/clstokes/aero/commands"
	"github.com/clstokes/aero/providers"
	"github.com/clstokes/aero/structs"
	"github.com/mitchellh/cli"
	"os"
)

const VERSION = "0.1.0"

type Factory func() structs.Provider

var fingerprinters = map[string]Factory{
	"amazon": providers.InitAmazon,
	"google": providers.InitGoogle,
}

var currentProvider structs.Provider

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	cli := &cli.CLI{
		Args:     os.Args[1:],
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
		provider := val()
		if provider.IsCurrentProvider() {
			return provider
		}
	}
	return nil
}
