package goBoostStruct

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Structure() {
	root := "."
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		depth := strings.Count(relativePath, string(os.PathSeparator))

		if d.IsDir() {
			if relativePath != "." {
				fmt.Printf("%s%s/\n", strings.Repeat("|  ", depth), d.Name())
			}
		} else {

			fmt.Printf("%s|__%s\n", strings.Repeat("|  ", depth), d.Name())
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error", err)
	}
}
