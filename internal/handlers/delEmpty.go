package handlers

import (
	"fmt"
	"os"
	"path/filepath"
)

func DelEmpty() {
	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Проверяем, что это директория
		if d.IsDir() {
			// Получаем список файлов в директории
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			// Если директория пуста, удаляем ее
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
