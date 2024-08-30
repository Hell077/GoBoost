package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "strings"
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

	fmt.Println("Checking for updates...")

	// Download the latest version of the executable
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Invoke-WebRequest -Uri %s -OutFile %s", releaseURL, tempFilePath))
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error downloading latest version: %v\n", err)
		return
	}

	// Replace the old executable with the new one
	if err := os.Rename(tempFilePath, executablePath); err != nil {
		fmt.Printf("Error replacing executable: %v\n", err)
		return
	}

	fmt.Println("Update successful. Please restart the program.")
}
