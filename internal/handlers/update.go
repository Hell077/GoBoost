package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func downloadNewVersion(releaseURL, tempFilePath string) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(releaseURL)
	if err != nil {
		return fmt.Errorf("error fetching the update file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: status code %d while fetching the update file", resp.StatusCode)
	}

	out, err := os.Create(tempFilePath)
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving update file: %v", err)
	}

	fmt.Println("New version downloaded successfully.")
	return nil
}

func terminateCurrentTerminal() {
	fmt.Println("Terminating the current terminal...")
	syscall.Exit(0)
}

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

	if err := downloadNewVersion(releaseURL, tempFilePath); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	updateScriptPath := filepath.Join(executableDir, "update_script.bat")
	cleanupScriptPath := filepath.Join(executableDir, "cleanup_script.bat")
	scriptContent := fmt.Sprintf(`@echo off
REM Start update in new terminal
start cmd /c "update.bat"

REM Close old terminal
exit
`, executablePath, tempFilePath)

	if err := os.WriteFile(updateScriptPath, []byte(scriptContent), 0755); err != nil {
		fmt.Printf("Error creating update script: %v\n", err)
		return
	}

	updateBatchContent := fmt.Sprintf(`@echo off
REM Update Script

echo Updating the application...
move /Y "%s" "%s"
echo Update complete!
REM Start clean script
start /b cmd /c "timeout /t 10 & %s"
`, tempFilePath, executablePath, cleanupScriptPath)

	if err := os.WriteFile(filepath.Join(executableDir, "update.bat"), []byte(updateBatchContent), 0755); err != nil {
		fmt.Printf("Error creating update batch script: %v\n", err)
		return
	}

	cleanupBatchContent := `@echo off
REM clean script

REM Delete temp files and script 
del "%~dp0goBoost_temp.exe"
del "%~dp0update.bat"
del "%~dp0update_script.bat"
echo Cleanup complete!
del cleanup_script.bat
`

	if err := os.WriteFile(cleanupScriptPath, []byte(cleanupBatchContent), 0755); err != nil {
		fmt.Printf("Error creating cleanup batch script: %v\n", err)
		return
	}
	cmd := exec.Command("cmd.exe", "/c", updateScriptPath)
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting update script: %v\n", err)
		return
	}
	terminateCurrentTerminal()
}
