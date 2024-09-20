package webapp

import (
	"github.com/Hell077/GoBoost/internal/handlers"
	"github.com/Hell077/GoBoost/internal/utils"
)

func CreateWebAppTemplate(projectDir, projectName string) {
	directories := []string{
		"cmd/app",
		"api",
		"web",
		"config",
		"migrations",
	}

	files := map[string]string{
		"cmd/app/main.go": `package main

func main() {
	// Start your web server here
}`,
		"api/api.go": `package api

// Define your API routes here
`,
		"config/config.yaml": `# Configuration file
`,
		"README.md": "# " + projectName + ` Web Application

## Overview

This is a web application project using Go.
`,
		".gitignore": handlers.Ignore(),
	}

	utils.CreateProjectStructure(projectDir, directories, files, projectName)
}
