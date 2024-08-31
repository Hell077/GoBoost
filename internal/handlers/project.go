package handlers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateProject() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter project name: ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading project name: %v\n", err)
		return
	}
	projectName = strings.TrimSpace(projectName)
	if projectName == "" {
		projectName = "default"
	}

	fmt.Print("Enter project path (or '.' for current directory): ")
	projectPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading project path: %v\n", err)
		return
	}
	projectPath = strings.TrimSpace(projectPath)
	if projectPath == "." {
		projectPath, err = os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
	}

	projectDir := filepath.Join(projectPath, projectName)
	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists\n", projectDir)
		return
	}

	fmt.Printf("Creating Go project: %s in %s\n", projectName, projectPath)

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


## Created by

[GoBoost](https://github.com/Hell077/GoBoost)`,
		".gitignore": `bin/
*.exe
*.log`,
	}
	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(projectDir, dir), os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}

	for filePath, content := range files {
		err := os.WriteFile(filepath.Join(projectDir, filePath), []byte(content), os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
	}

	if err := os.Chdir(projectDir); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}

	if _, err := exec.Command("go", "mod", "init", projectName).Output(); err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
		return
	}

	fmt.Println("Go project created successfully!")
}
