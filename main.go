package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func listFiles(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s:\n", dir) // Print the directory name as a header
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

func main() {
	dirs := os.Args[1:] // Get directories from command-line arguments
	if len(dirs) == 0 {
		dirs = append(dirs, ".") // Default to current directory if none provided
	}

	for _, dir := range dirs {
		// Get absolute path for clarity
		absDir, err := filepath.Abs(dir)
		if err != nil {
			fmt.Printf("Error getting absolute path: %v\n", err)
			continue
		}

		// List files in the directory
		if err := listFiles(absDir); err != nil {
			fmt.Printf("Error reading directory %s: %v\n", absDir, err)
		}
	}
}
