To extend the CLI with a flag that allows copying the compiled executable to a remote host using SSH, we can use the `flag` package to handle command-line arguments and the `os/exec` package to run the `scp` command for secure copying.

### Extended Code Explanation
1. The code adds a new flag `--remote` (or `-r`) to specify the remote host in the format `user@host:/path/to/destination`.
2. After compiling the executables, if the remote flag is provided, the CLI will use the `scp` command to copy the compiled files to the specified remote location.
3. It uses `exec.Command` to run the `scp` command.

Here's the updated code:

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// compileProgram compiles the given source file for the specified OS and architecture.
func compileProgram(goos, goarch, sourceFile string) (string, error) {
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

	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Return the name of the compiled output
	return outputName, nil
}

// copyToRemote copies the compiled executable to a remote host using scp.
func copyToRemote(remote, file string) error {
	fmt.Printf("Copying %s to %s...\n", file, remote)
	// Split remote into user@host:/path
	parts := strings.Split(remote, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid remote format, expected user@host:/path")
	}

	// Run the scp command
	cmd := exec.Command("scp", file, remote)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Define command-line flags
	remote := flag.String("remote", "", "Remote destination in the format user@host:/path/to/destination")
	flag.Parse()

	// Ensure the source file is provided as an argument
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run build.go <source_file.go> [--remote user@host:/path/to/destination]")
		os.Exit(1)
	}

	sourceFile := flag.Arg(0)

	// Compile for Linux
	linuxOutput, err := compileProgram("linux", "amd64", sourceFile)
	if err != nil {
		fmt.Printf("Failed to compile for Linux: %v\n", err)
	}

	// Compile for Windows
	windowsOutput, err := compileProgram("windows", "amd64", sourceFile)
	if err != nil {
		fmt.Printf("Failed to compile for Windows: %v\n", err)
	}

	// Copy to remote host if specified
	if *remote != "" {
		// Copy the Linux executable
		if linuxOutput != "" {
			if err := copyToRemote(*remote, linuxOutput); err != nil {
				fmt.Printf("Failed to copy Linux executable: %v\n", err)
			}
		}

		// Copy the Windows executable
		if windowsOutput != "" {
			if err := copyToRemote(*remote, windowsOutput); err != nil {
				fmt.Printf("Failed to copy Windows executable: %v\n", err)
			}
		}
	}

	fmt.Println("Compilation completed.")
}
```

### How to Use the Extended CLI
1. Save the code as `build.go`.
2. Compile a Go source file using the following command:

   ```bash
   go run build.go path/to/your/source.go
   ```

3. To compile and copy the executable to a remote host, use:

   ```bash
   go run build.go path/to/your/source.go --remote user@host:/path/to/destination
   ```

### Key Points
- **`--remote` Flag**: Use the `--remote` flag to specify the destination for the `scp` command.
- **Error Handling**: The CLI checks for invalid remote formats and errors during copying.
- **Dependencies**: Ensure SSH access is configured (e.g., public key authentication) for the remote host to avoid password prompts during `scp`.

This approach provides a streamlined way to compile Go programs and deploy them to remote environments directly! Let me know if you need any more tweaks.