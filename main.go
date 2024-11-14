/*
The program lists files in the current directory if no argument is given,
or in a specified directory if one is provided as a command-line argument.

	It ignores any errors while reading directory contents, which could lead to
	issues if the directory does not exist.

	code developed by Bhamare Sanket
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		files, _ := os.ReadDir("./")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		files, _ := os.ReadDir(os.Args[1])
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
}
