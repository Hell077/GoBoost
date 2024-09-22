package main

import (
	"fmt"
	"github.com/Hell077/GoBoost/internal/goBoostBuild"
	"github.com/Hell077/GoBoost/internal/goBoostStruct"
	"github.com/Hell077/GoBoost/internal/handlers"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		handlers.CreateProject()
		return
	}

	if len(args[0]) == 0 {
		fmt.Println("Error: Empty flag detected. Please provide a valid flag.")
		return
	}

	switch strings.ToLower(args[0]) {
	case "--help", "--h":
		handlers.ShowHelp()
		return
	case "--update", "--u":
		handlers.UpdateProgram()
		return
	case "--build", "--b":
		goBoostBuild.Build()
		return
	case "--struct", "--s":
		goBoostStruct.Structure()
		return
	case "create":
		handlers.CreateProject()
		return
	default:
		fmt.Printf("Error: Invalid flag '%s'. Use --help or -h for available commands.\n", args[0])
		return
	}
}
