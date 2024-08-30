package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const version = "1.1.0"

func ReadVersionFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // Файл не существует, версия не задана
		}
		return "", err
	}
	return string(data), nil
}

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
	versionFilePath := filepath.Join(executableDir, "current_version.txt")

	fmt.Println("Checking for updates...")

	// Read the current version from file
	currentVersion, err := ReadVersionFromFile(versionFilePath)
	if err != nil {
		fmt.Printf("Error reading current version from file: %v\n", err)
		return
	}

	// Compare with the new version
	if currentVersion == version {
		fmt.Println("The program is already up-to-date.")
		return
	}

	// Write the new version to file
	if err := WriteVersionToFile(versionFilePath, version); err != nil {
		fmt.Printf("Error writing version to file: %v\n", err)
		return
	}

	// Download the latest version
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Invoke-WebRequest -Uri %s -OutFile %s", releaseURL, tempFilePath))
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Error downloading latest version: %v\nOutput: %s\n", err, output)
		return
	}

	// Create an update script to replace the current executable
	scriptContent := fmt.Sprintf(
		`@echo off
echo %%date%% %%time%%: Script started > update.log
timeout /t 5 /nobreak > NUL
echo %%date%% %%time%%: Deleting old executable >> update.log
if exist "%s" (
    del "%s" >> update.log 2>&1
    if errorlevel 1 echo Failed to delete goboost.exe >> update.log
) else (
    echo goboost.exe does not exist >> update.log
)
echo %%date%% %%time%%: Moving new executable >> update.log
move /Y "%s" "%s" >> update.log 2>&1
if errorlevel 1 echo Failed to move goBoost_temp.exe >> update.log
echo %%date%% %%time%%: Script completed >> update.log
`,
		executablePath, executablePath, tempFilePath, executablePath,
	)
	if err := os.WriteFile(updateScriptPath, []byte(scriptContent), 0755); err != nil {
		fmt.Printf("Error creating update script: %v\n", err)
		return
	}

	// Run the update script in a separate process
	cmd = exec.Command("cmd.exe", "/c", updateScriptPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Error running update script: %v\nOutput: %s\n", err, output)
		return
	}

	// Optionally, you can delete the update script after execution
	if err := os.Remove(updateScriptPath); err != nil {
		fmt.Printf("Error removing update script: %v\n", err)
		return
	}

	// Delete the temporary file after update
	if err := os.Remove(tempFilePath); err != nil {
		fmt.Printf("Error removing temporary file: %v\n", err)
		return
	}

	// Give the system some time to complete the update
	time.Sleep(2 * time.Second)

	fmt.Println("Update initiated. Please restart the program to complete the update.")
}
