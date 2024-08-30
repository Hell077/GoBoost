# Go Project Initializer

A simple tool for initializing a new Go project with a predefined structure.

# Releases
**You can download pre-built binaries from the Releases section. Each release includes:**
```
Windows: goBoost.exe
```

## First Installation

1. **Clone this repository:**

   ```bash
   git clone https://github.com/Hell077/GoBoost.git
   ```
2. **Navigate to the directory and build the executable**
   ```bash
    cd path-to-project
    go build -o goBoost.exe
   ```
3. **Move the executable to a directory in your PATH (Optional)**
    ```bash
    Copy-Item -Path .\goBoost.exe -Destination "C:\Program Files\goBoost.exe"
   ```

## Second Installation

1.  **Download the .exe file from the Releases page.**
2. **Move the extracted files to a directory of your choice.
  <br/> For example, you might place it in: 
    ```bash
   C:\Program Files\GoBoost.**
   ```
3. Add GoBoost to PATH:
    - Open the Start Menu and search for "Environment Variables."<br/>
    - Select "Edit the system environment variables."<br/>
    - In the System Properties window, click on the "Environment Variables" button.<br/>
    - In the Environment Variables window, locate the "Path" variable in the "System variables" section and click "Edit."<br/>
    - Click "New" and add the path to the directory where you moved GoBoost (e.g., C:\Program Files\GoBoost).<br/>
    - Click "OK" to close all open windows and apply the changes.<br/>
4. **Usage**
   - ``in cmd or ide or another terminal on windows`` <br/>
   - ``follow the instructions inside the program``
  ```bash
   goBoost
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



