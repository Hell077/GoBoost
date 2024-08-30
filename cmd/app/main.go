package main

import (
	"github.com/Hell077/GoBoost/internal/handlers"
	"os"
	"strings"
)

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
		default:
			if args[0] == "create" {
				handlers.CreateProject()
				return
			}
		}
	}

	handlers.CreateProject()
}
