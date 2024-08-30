package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// UpdateProgram checks if the binary is up-to-date and updates it if necessary.
func UpdateProgram() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		return
	}

	executableDir := filepath.Dir(executablePath)
	releaseURL := "https://github.com/Hell077/GoBoost/releases/latest/download/goBoost.exe"
	tempFilePath := filepath.Join(executableDir, "goBoost_temp.exe")
	updateScriptPath := filepath.Join(executableDir, "update.bat")

	fmt.Println("Checking for updates...")

	// Download the latest version of the executable
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Invoke-WebRequest -Uri %s -OutFile %s", releaseURL, tempFilePath))
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error downloading latest version: %v\n", err)
		return
	}

	// Create an update batch script
	scriptContent := fmt.Sprintf(
		`@echo off
		timeout /t 5 /nobreak > NUL
		del "%s"
		move "%s" "%s"
	`,
		executablePath, tempFilePath, executablePath,
	)
	if err := os.WriteFile(updateScriptPath, []byte(scriptContent), 0755); err != nil {
		fmt.Printf("Error creating update script: %v\n", err)
		return
	}

	// Execute the update script
	cmd = exec.Command(updateScriptPath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running update script: %v\n", err)
		return
	}

	// Optionally, you can delete the update script after execution
	if err := os.Remove(updateScriptPath); err != nil {
		fmt.Printf("Error removing update script: %v\n", err)
		return
	}

	fmt.Println("Update successful. Please restart the program.")
}
