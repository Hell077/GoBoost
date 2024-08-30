package handlers

import (
	"fmt"
)

// Строка с данными для отображения
const data = `
  - \cmd\
  - \internal\
  - \pkg\
  - \README.md\
  - \.gitignore\
`

// READMEContent содержит содержимое README.md в виде строки с включением переменной data.
const READMEContent = `
# GoBoost

## 🚀 New Features
- **GoBoost**: A new command-line tool to streamline the setup of Go projects with a predefined structure.
- **Automatic Project Setup**: Creates essential directories and files, including:
` + data + `
## 🛠️ Improvements
- **Enhanced Error Handling**: Improved responses to scenarios involving existing directories and invalid inputs.
- **User Prompts**: Refined prompts for clearer user interaction and accurate input handling.

## 🔧 Fixes
- **Path Handling**: Fixed issues with handling relative and absolute paths during project creation.
- **Directory Checks**: Resolved bugs related to checking directory existence to prevent overwriting existing projects.

## 📜 Instructions
1. **Download** the appropriate binary for your operating system from the links above.
2. **Run** the executable from your command line or terminal.
3. **Follow** the prompts to provide your project name and the desired path.
4. **Verify** that the project directory contains the initialized Go project structure.

## 🔗 Links
- [Documentation](https://github.com/Hell077/GoBoost/blob/main/README.MD)
- [Source Code](https://github.com/Hell077/GoBoost)

## 🔍 Known Issues
- **None**

## 📅 Date
- **August 31, 2024**
`

func ShowHelp() {
	fmt.Println(READMEContent)
}
