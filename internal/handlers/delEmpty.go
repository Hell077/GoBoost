package handlers

import (
	"fmt"
	"os"
	"path/filepath"
)

func DelEmpty() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			if len(entries) == 0 {
				fmt.Printf("Deleting empty directory: %s\n", path)
				if err := os.Remove(path); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
