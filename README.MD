# Go Project Initializer

A simple tool for initializing a new Go project with a predefined structure.

# Releases
**You can download pre-built binaries from the Releases section. Each release includes:**
```
Windows: goBoost.exe
```

## Installation

1. **Clone this repository:**

   ```bash
   git clone https://github.com/Hell077/GoBoost.git
   ```
2. **Navigate to the directory and build the executable**
   ```bash
    cd go-project-initializer
    go build -o go-init.exe
   ```
3. **Move the executable to a directory in your PATH (Optional)**
    ```bash
    Copy-Item -Path .\goBoost.exe -Destination "C:\Program Files\goBoost.exe"
   ```



**Usage**

  ```bash
   goBoost.exe
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



