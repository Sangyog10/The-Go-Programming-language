package main

import (
	"fmt"
	"os"
	// "strings"
	// "time"
)

func main() {
	// fmt.Println("Command line arguments path:",os.Args[0]) //gives the path of program
	fmt.Println("Command line arguments:", os.Args[1:])
	user_input()

	// $go run index.go hello world
	// output: Command line arguments: [hello world]

}

//simulating the echo command

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(os.Args[i] + " ")
	}
	fmt.Println()
}

//Alternative approach

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Print(arg + " ")
	}
	fmt.Println()

	// _ is used to ignore the initialization
	// Args[1:]=> gives a subslice after 1 to the end(slice)
	// slices are dynamically sized collections. its like js array or c++ vectors
}

func main() {
	for index, arg := range os.Args {
		fmt.Println("Argument %d: %s\n", index, arg)
	}
	//$go run basic.go one two three
}

// time comparision between strings.Join and concat(+=)
func main() {
	//for +=
	start1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Time taken for +=:", time.Since(start1))

	//for join
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Time taken with `strings.Join`: %v\n", time.Since(start2))
}


//taking input from user
func user_input(){
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is: ", name)
}