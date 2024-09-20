package microservices

import (
	"github.com/Hell077/GoBoost/internal/utils"
)

func CreateMicroservicesTemplate(projectDir, projectName string) {
	directories := []string{
		"cmd/app",
		"services/auth",
		"services/payment",
		"gateway",
		"proto",
		"deploy",
	}

	files := map[string]string{
		"cmd/app/main.go": `package main

func main() {
	// Start your microservices here
}`,
		"gateway/gateway.go": `package gateway

// Implement your API gateway here
`,
		"proto/service.proto": `syntax = "proto3";

// Define your gRPC services here
`,
		"README.md": "# " + projectName + ` Microservices

## Overview

This is a microservices project using Go.
`,
		".gitignore": utils.Ignore(),
	}

	utils.CreateProjectStructure(projectDir, directories, files, projectName)
}
