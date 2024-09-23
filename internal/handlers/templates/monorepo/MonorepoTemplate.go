package monorepo

import (
	"github.com/Hell077/GoBoost/internal/utils"
)

func CreateMonorepoTemplate(projectDir, projectName string) {
	directories := []string{
		"apps/app1",
		"apps/app2",
		"libs/common",
		"tools",
		"infra",
	}

	files := map[string]string{
		"apps/app1/main.go": `package main

func main() {
	// Start app1 here
}`,
		"libs/common/utils.go": `package common

// Common utility functions
`,
		"tools/tool.go": `package tools

// Implement your tools here
`,
		"infra/terraform.tf": `# Infrastructure configuration
`,
		"README.md": "# " + projectName + ` Monorepo

## Overview

This is a monorepo project using Go.
`,
		".gitignore":     utils.Ignore(),
		".gitattributes": utils.Attribute(),
	}

	utils.CreateProjectStructure(projectDir, directories, files, projectName)
}
