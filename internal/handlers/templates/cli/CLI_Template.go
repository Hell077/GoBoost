package cli

import "github.com/Hell077/GoBoost/internal/utils"

func CreateCLIAppTemplate(projectDir, projectName string) {
	directories := []string{
		"cmd/app",
		"internal/cli",
		"scripts",
		"assets",
	}

	files := map[string]string{
		"cmd/app/main.go": `package main

func main() {
	// Start your CLI application here
}`,
		"internal/cli/cli.go": `package cli

// Define your CLI commands here
`,
		"scripts/build.sh": `#!/bin/bash
# Build script
`,
		"README.md": "# " + projectName + ` CLI Application

## Overview

This is a CLI application project using Go.
`,
		".gitignore": `bin/
*.exe
*.log`,
	}

	utils.CreateProjectStructure(projectDir, directories, files, projectName)
}
