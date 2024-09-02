package goBoostBuild

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var selected string = ""

func Build() {
	fmt.Println("Enter the path to your main file (e.g., cmd/app/main.go) or use (.) to search for main.go in the project:")

	_, err := fmt.Scanln(&selected)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if selected == "" || selected == "." {
		selected, err = findMainGo(".")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	cmd := exec.Command("go", "build", "-ldflags", "-w -s", selected)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Build failed:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Build successful!")
}

func findMainGo(root string) (string, error) {
	var mainGoPath string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "main.go" {
			mainGoPath = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if mainGoPath == "" {
		return "", fmt.Errorf("main.go not found in project")
	}

	return mainGoPath, nil
}
