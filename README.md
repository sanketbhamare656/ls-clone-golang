# Directory Lister

A simple Go program that lists all files in a specified directory. If no directory is specified, it defaults to listing files in the current directory.

## Usage

### Clone the Repository

```bash
git clone https://github.com/sanketbhamare656/ls-clone-golang.git
cd ls-clone-golang

Running the Program
To use this program, you need Go installed.

Run the program in one of the following ways:

List Files in the Current Directory

If you run the program without any arguments, it will list all files in the current directory.

bash
Copy code
go run main.go
List Files in a Specified Directory

To list files in a specific directory, provide the directory path as an argument:

bash
Copy code
go run main.go <directory-path>
Replace <directory-path> with the path to the directory you want to list.

Code Explanation
Here’s a line-by-line explanation of how the code works:

The program checks if a directory path is provided as an argument.
If no path is provided, it defaults to listing files in the current directory.
If a path is provided, it lists the contents of the specified directory.
The program uses the os and fmt packages in Go:
os.ReadDir to read directory contents.
fmt.Println to print file names to the console.
Example
bash
Copy code
go run main.go /path/to/your/directory
Output:

plaintext
Copy code
file1.txt
file2.txt
file3.go
Error Handling
The program ignores errors during directory reading. In case the directory does not exist or cannot be accessed, it may not provide useful feedback. You can modify the code to handle errors if needed.

License
This project is licensed under the MIT License. See the LICENSE file for details.

rust
Copy code

This `README.md` provides setup, usage instructions, and example usage for your program, making it suitable for sharing on GitHub.





You said:

