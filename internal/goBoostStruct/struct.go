package goBoostStruct

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Structure() {
	root := "." // текущая директория
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Получаем относительный путь для корректного отображения вложенности
		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Разделяем путь на директории
		depth := strings.Count(relativePath, string(os.PathSeparator))

		// Если это директория, выводим ее название без символов |__ для файлов
		if d.IsDir() {
			if relativePath != "." { // игнорируем текущую директорию
				fmt.Printf("%s%s/\n", strings.Repeat("|  ", depth), d.Name())
			}
		} else {
			// Для файлов используем |__ перед именем файла
			fmt.Printf("%s|__%s\n", strings.Repeat("|  ", depth), d.Name())
		}

		return nil
	})

	if err != nil {
		fmt.Println("Ошибка при обходе файловой системы:", err)
	}
}
