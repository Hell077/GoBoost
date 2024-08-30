# GoBoost

GoBoost is a simple tool for initializing a new Go project with a predefined structure. It also supports updating itself to the latest version with a single command.

## Releases
**You can download pre-built binaries from the Releases section. Each release includes:**


## First Installation

 **Clone this repository:**
   ```bash
   git clone https://github.com/Hell077/GoBoost.git
   ```
   ```bash
   cd path-to-project
   go build -o goBoost.exe
   ```
   ```bash 
   Copy-Item -Path .\goBoost.exe -Destination "C:\Program Files\goBoost.exe"
   ```
## Second Installation
1. Download the .exe file from the Releases page.

2. Move the downloaded file to a directory of your choice. For example, you might place it in:

   ````
   C:\Program Files\GoBoost
   ````
3. Add GoBoost to PATH:

- Open the Start Menu and search for "Environment Variables."
- Select "Edit the system environment variables."
- In the System Properties window, click on the "Environment Variables" button.
- In the Environment Variables window, locate the "Path" variable in the "System variables" section and click "Edit."
- Click "New" and add the path to the directory where you moved GoBoost (e.g., C:\Program Files\GoBoost).
- Click "OK" to close all open windows and apply the changes.

## Usage

- Open a command prompt, IDE, or another terminal on Windows.
- Follow the instructions inside the program.

## Commands
-  goBoost --help or goBoost -h: Displays the content of README.md.
-  goBoost --update: Checks for updates and replaces the current executable with the latest version

## Project Structure

````
<project-name>
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   └── server/
│       └── server.go
├── pkg/
│   └── utils/
│       └── utils.go
├── README.md
└── .gitignore

````