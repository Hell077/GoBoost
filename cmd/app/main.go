package main

import (
	"fmt"
	"github.com/Hell077/GoBoost/internal/handlers"
	"os"
	"strings"
)

const version = "1.1.0"

func printVersion() {
	fmt.Printf("Current version: %s\n", version)
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		switch strings.ToLower(args[0]) {
		case "--help", "-h":
			handlers.ShowHelp()
			return
		case "--update":
			handlers.UpdateProgram()
			return
		case "--version":
			printVersion()
			return
		default:
			if args[0] == "create" {
				handlers.CreateProject()
				return
			}
		}
	}

	handlers.CreateProject()
}
