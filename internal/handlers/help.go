package handlers

import (
	"bufio"
	"fmt"
	"os"
)

// ShowHelp displays the contents of the README.md file.
func ShowHelp() {
	readmeFile := "README.md"

	file, err := os.Open(readmeFile)
	if err != nil {
		fmt.Printf("Error opening README.md: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading README.md: %v\n", err)
	}
}
