package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateProjectStructure(projectDir string, directories []string, files map[string]string, projectName string) {
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
