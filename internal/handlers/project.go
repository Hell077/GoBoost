package handlers

import (
	"bufio"
	"fmt"
	"github.com/Hell077/GoBoost/internal/handlers/templates/apiGateway"
	"github.com/Hell077/GoBoost/internal/handlers/templates/cli"
	"github.com/Hell077/GoBoost/internal/handlers/templates/default"
	"github.com/Hell077/GoBoost/internal/handlers/templates/microservices"
	"github.com/Hell077/GoBoost/internal/handlers/templates/monorepo"
	"github.com/Hell077/GoBoost/internal/handlers/templates/webapp"
	"github.com/manifoldco/promptui"
	"os"
	"path/filepath"
	"strings"
)

func CreateProject() {
	prompt := promptui.Select{
		Label: "Select Project Template",
		Items: []string{"Default", "Web App Template", "Microservices Template", "CLI App Template", "Monorepo Template", "Api Gateway Template"},
	}

	_, template, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error selecting template: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter project name: ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading project name: %v\n", err)
		return
	}
	projectName = strings.TrimSpace(projectName)
	if projectName == "" {
		projectName = "default"
	}

	fmt.Print("Enter project path (or '.' for current directory): ")
	projectPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading project path: %v\n", err)
		return
	}
	projectPath = strings.TrimSpace(projectPath)
	if projectPath == "." {
		projectPath, err = os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
	}

	projectDir := filepath.Join(projectPath, projectName)
	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists\n", projectDir)
		return
	}

	fmt.Printf("Creating Go project: %s in %s using template: %s\n", projectName, projectPath, template)

	switch template {
	case "Default":
		_default.CreateDefaultProject(projectDir, projectName)
	case "Web App Template":
		webapp.CreateWebAppTemplate(projectDir, projectName)
	case "Microservices Template":
		microservices.CreateMicroservicesTemplate(projectDir, projectName)
	case "CLI App Template":
		cli.CreateCLIAppTemplate(projectDir, projectName)
	case "Monorepo Template":
		monorepo.CreateMonorepoTemplate(projectDir, projectName)
	case "Api Gateway Template":
		apiGateway.CreateApiGatewayTemplate(projectDir, projectName)
	}
}
