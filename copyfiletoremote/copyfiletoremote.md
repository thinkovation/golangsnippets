Here's a Go CLI application that copies a file to a remote host using SCP, with the copy functionality encapsulated in a separate module that can be reused elsewhere. The implementation will include:

1. **Main CLI application (`main.go`)**: This will handle command-line arguments and call the copy function.
2. **Copy module (`copyutil/copyutil.go`)**: This module will contain the logic for copying files to a remote host using SCP.

### Step 1: Create the Copy Module (`copyutil/copyutil.go`)

First, let's create the `copyutil` module that handles the SCP functionality. This module provides a reusable function `CopyFileToRemote` that can be used by other applications.

```go
// File: copyutil/copyutil.go
package copyutil

import (
	"fmt"
	"os/exec"
	"strings"
)

// CopyFileToRemote copies a file to a remote host using SCP.
// The remote parameter should be in the format user@host:/path/to/destination.
// The function returns an error if the copy fails.
func CopyFileToRemote(remote, file string) error {
	fmt.Printf("Copying %s to %s...\n", file, remote)
	
	// Split remote into user@host:/path to validate the format
	parts := strings.Split(remote, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid remote format, expected user@host:/path")
	}

	// Prepare the SCP command
	cmd := exec.Command("scp", file, remote)
	cmd.Stdout = nil  // Hide command output
	cmd.Stderr = nil  // Hide error output

	// Execute the SCP command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	fmt.Printf("Successfully copied %s to %s\n", file, remote)
	return nil
}
```

### Step 2: Create the CLI Application (`main.go`)

Next, let's create the main CLI application that uses the `copyutil` module to perform the file copy.

```go
// File: main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yourusername/projectname/copyutil" // Change this import path to match your module structure
)

func main() {
	// Define command-line flags
	remote := flag.String("remote", "", "Remote destination in the format user@host:/path/to/destination")
	file := flag.String("file", "", "Path to the local file to copy")
	flag.Parse()

	// Check if the remote and file arguments are provided
	if *remote == "" || *file == "" {
		fmt.Println("Usage: go run main.go --file <file> --remote <user@host:/path/to/destination>")
		os.Exit(1)
	}

	// Perform the file copy using the copyutil module
	err := copyutil.CopyFileToRemote(*remote, *file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
```

### How to Set Up and Use the CLI

1. **Set Up the Project Structure**: Organize your project as follows:

   ```
   projectname/
   ├── main.go
   └── copyutil/
       └── copyutil.go
   ```

2. **Replace Import Path**: Make sure to replace the import path `github.com/yourusername/projectname/copyutil` in `main.go` with the actual path where your project resides.

3. **Run the CLI**: Use the following command to copy a file to a remote host:

   ```bash
   go run main.go --file /path/to/local/file --remote user@host:/path/to/destination
   ```

### Key Features

- **Reusable Module**: The `copyutil` module encapsulates the copy functionality, allowing you to use `CopyFileToRemote` in other projects.
- **Error Handling**: The module checks the format of the remote argument and handles any errors during the SCP process.
- **CLI Flexibility**: The CLI accepts command-line arguments for both the file and the remote destination, making it versatile and easy to use.

This setup provides a clean separation between the CLI logic and the reusable copy functionality, making it easier to maintain and extend in the future. Let me know if you have any questions or need further adjustments!