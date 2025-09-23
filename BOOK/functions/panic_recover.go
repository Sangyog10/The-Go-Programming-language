package main

import (
	"fmt"
)

func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    fmt.Println("Starting risky function")
    panic("Something went wrong!")
    fmt.Println("This line will not execute")
}

func main() {
    riskyFunction()
    fmt.Println("Program continues after panic recovery")
}

// panic triggers a runtime error, stopping normal execution and unwinding the stack until all deferred calls are executed. This helps clean up resources.
// recover is only useful inside deferred functions. It stops the panic and allows the program to continue execution.