# Go Project Initializer

A simple tool for initializing a new Go project with a predefined structure.

## Installation

1. **Clone this repository:**

   ```bash
   git clone https://github.com/yourusername/go-project-initializer.git
   ```
2. **Navigate to the directory and build the executable**
   ```bash
    cd go-project-initializer
    go build -o go-init.exe
   ```
3. **Move the executable to a directory in your PATH (Optional)**
    ```bash
    Copy-Item -Path .\go-init.exe -Destination "C:\Program Files\go-init.exe"
   ```

**Usage**

  ```bash
   go-init.exe
   ```

**Project Structure**

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



