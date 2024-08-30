package handlers

import (
	"os"
)

func WriteVersionToFile(filePath, version string) error {
	return os.WriteFile(filePath, []byte(version), 0644)
}
