// How this is a script the main can cause an warning in some IDEs, but it is a valid Go program.
// The main function is the entry point of the program, and it is where the execution starts. In this case,
// the main function is used to update the go.work file with all the modules found in the current directory and its subdirectories.
package main

// For build a script without a module, you need to use this command
// GO111MODULE=off go build -o update-go-work update-go-work.go

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var dirs []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "go.mod" {
			dirs = append(dirs, filepath.Dir(path))
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		os.Exit(1)
	}

	if len(dirs) == 0 {
		fmt.Println("No go.mod files found in the current directory or its subdirectories.")
	}

	// Decide between init or use
	var cmd *exec.Cmd

	if _, err := os.Stat("go.work"); os.IsNotExist(err) {
		// If dont exist, initialize the workspace with all modules
		args := append([]string{"work", "init"}, dirs...)
		cmd = exec.Command("go", args...) // go work init <dirs...>
		fmt.Println("Updating go work: %v\n", args)
	} else {
		// If already exists, use the workspace with all modules
		args := append([]string{"work", "use"}, dirs...)
		cmd = exec.Command("go", args...) // go work use <dirs...>
		fmt.Println("Updating go work: %v\n", args)
	}

	// The command's standard output and standard error are set to the program's standard output and standard error.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Go allows a short variable declaration before the condition in an if statement.
	// The variable exists only inside the if/else scope.
	//
	// Here, cmd.Run() is executed and its returned error is stored in err.
	// If err is not nil, the command failed and the error-handling block is executed.
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running go command:", err)
		os.Exit(1)
	}

	fmt.Println("go.work file updated successfully.")
}
