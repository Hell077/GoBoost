package handlers

import (
	"fmt"
	"os/exec"
	"runtime"
)

const docURL = "https://github.com/Hell077/GoBoost/blob/main/README.MD"

func ShowHelp() {
	// Определяем, какую команду использовать для открытия браузера в зависимости от ОС
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", docURL)
	case "darwin":
		cmd = exec.Command("open", docURL)
	case "linux":
		cmd = exec.Command("xdg-open", docURL)
	default:
		fmt.Println("Unsupported platform")
		return
	}

	// Пытаемся выполнить команду
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to open documentation:", err)
		return
	}

	fmt.Println("Documentation opened in your browser.")
}
