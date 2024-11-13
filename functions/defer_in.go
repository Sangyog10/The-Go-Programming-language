package main

import (
"fmt"
"os"
)

// Deferred Function Calls(delay the execution of func until the surrounding func return)
func main(){
	file ,err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error in file creation ,", err)
		return
	}
	defer file.Close() // Ensures file is closed when `main` exits
	fmt.Fprintln(file, "Hello, Go!")
    fmt.Println("File written successfully.")
} 

// defer file.Close() ensures that file.Close() is called after main finishes, even if an error occurs.
// defer is often used for cleanup tasks, making the code clearer and preventing resource leaks.
// Deferred calls execute in last-in, first-out order. If you defer multiple functions, they’ll execute in reverse order of their defer statements.