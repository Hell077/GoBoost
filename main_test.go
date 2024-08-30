package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Helper function to run the main application with custom inputs
func runMain(inputs string) (string, error) {
	cmd := exec.Command("go", "run", "main.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Simulate input by providing it to the standard input
	cmd.Stdin = strings.NewReader(inputs)
	err := cmd.Run()

	return out.String(), err
}

// Test the project creation functionality
func TestProjectCreation(t *testing.T) {
	// Define the project name and path for the test
	projectName := "testproject"
	projectPath := "."

	// Provide input to the main function
	inputs := projectName + "\n" + projectPath + "\n"

	// Run the main application
	output, err := runMain(inputs)
	if err != nil {
		t.Fatalf("Error running main application: %v", err)
	}

	// Check if the output contains the expected messages
	if !strings.Contains(output, "Creating Go project: testproject in") {
		t.Errorf("Output did not contain expected message: %s", output)
	}

	// Check if the project directory was created
	projectDir := filepath.Join(projectPath, projectName)
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		t.Errorf("Project directory %s does not exist", projectDir)
	}

	// Clean up: remove the test project directory
	err = os.RemoveAll(projectDir)
	if err != nil {
		t.Errorf("Error removing project directory: %v", err)
	}
}

// Test path handling
func TestPathHandling(t *testing.T) {
	// Define the project name and path for the test
	projectName := "testprojectpath"
	projectPath := "./"

	// Provide input to the main function
	inputs := projectName + "\n" + projectPath + "\n"

	// Run the main application
	output, err := runMain(inputs)
	if err != nil {
		t.Fatalf("Error running main application: %v", err)
	}

	// Check if the output contains the expected messages
	if !strings.Contains(output, "Creating Go project: testprojectpath in") {
		t.Errorf("Output did not contain expected message: %s", output)
	}

	// Check if the project directory was created
	projectDir := filepath.Join(projectPath, projectName)
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		t.Errorf("Project directory %s does not exist", projectDir)
	}

	// Clean up: remove the test project directory
	err = os.RemoveAll(projectDir)
	if err != nil {
		t.Errorf("Error removing project directory: %v", err)
	}
}
