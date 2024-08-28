Certainly! Below is a simple Go CLI program that compiles a Go program into executables for both Windows and Linux. The CLI will use `go build` command with the appropriate environment variables set to target different operating systems.

### Code Explanation
- The CLI accepts the path of the Go source file you want to compile.
- It uses `exec.Command` to run `go build` with the appropriate `GOOS` and `GOARCH` environment variables.
- The output executables are named based on the target operating system.

Here's the Go CLI code:

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// compileProgram compiles the given source file for the specified OS and architecture.
func compileProgram(goos, goarch, sourceFile string) error {
	// Get the source file name without extension for naming the executable
	sourceName := filepath.Base(sourceFile)
	outputName := sourceName[:len(sourceName)-len(filepath.Ext(sourceName))]

	// Append .exe for Windows executable
	if goos == "windows" {
		outputName += ".exe"
	}

	// Set the environment variables for cross-compilation
	cmd := exec.Command("go", "build", "-o", outputName, sourceFile)
	cmd.Env = append(os.Environ(),
		"GOOS="+goos,
		"GOARCH="+goarch,
	)

	// Run the build command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Compiling %s for %s/%s...\n", sourceFile, goos, goarch)

	return cmd.Run()
}

func main() {
	// Ensure the source file is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <source_file.go>")
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Compile for Linux
	if err := compileProgram("linux", "amd64", sourceFile); err != nil {
		fmt.Printf("Failed to compile for Linux: %v\n", err)
	}

	// Compile for Windows
	if err := compileProgram("windows", "amd64", sourceFile); err != nil {
		fmt.Printf("Failed to compile for Windows: %v\n", err)
	}

	fmt.Println("Compilation completed.")
}
```

### How to Use
1. Save the above code in a file, say `build.go`.
2. Run the CLI with the path to the Go source file you want to compile:

   ```bash
   go run build.go path/to/your/source.go
   ```

3. The program will create two executables, one for Linux and one for Windows, in the current directory.

### Key Points
- **GOOS** and **GOARCH** environment variables specify the target OS (`linux`, `windows`) and architecture (`amd64`).
- This script only targets `amd64` architecture; you can extend it for other architectures if needed.
- Error handling is included to display compilation errors, if any.

Let me know if you need further adjustments!