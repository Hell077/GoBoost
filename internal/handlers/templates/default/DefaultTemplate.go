package _default

import (
	"github.com/Hell077/GoBoost/internal/handlers"
	"github.com/Hell077/GoBoost/internal/utils"
)

func CreateDefaultProject(projectDir, projectName string) {
	directories := []string{
		"cmd/app",
		"internal/server",
		"pkg/utils",
	}

	files := map[string]string{
		"cmd/app/main.go": `package main

func main() {
	// TODO: Add your code here
}`,
		"internal/server/server.go": `package server

// TODO: Implement your server logic here
`,
		"pkg/utils/utils.go": `package utils

// TODO: Add utility functions here
`,
		"README.md": "# " + projectName + `

## Overview

This is a Go project named ` + projectName + `.
`,
		".gitignore": handlers.Ignore(),
	}
	utils.CreateProjectStructure(projectDir, directories, files, projectName)
}
