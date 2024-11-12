package main

import (
    "fmt" 
	"os"
)


// func main(){
// 	// fmt.Println("Command line arguments path:",os.Args[0]) //gives the path of program
// 	fmt.Println("Command line arguments:",os.Args[1:]) 
	
// 	// $go run index.go hello world
// 	// output: Command line arguments: [hello world]
	
// }
 

//simulating the echo command 

// func main(){
// 	for i := 1; i<len(os.Args); i++{
// 		fmt.Print(os.Args[i] + " ")
// 	}
// 	fmt.Println()
// }


//Alternative approach

func main(){
	for _,  arg := range os.Args[1:]{
		fmt.Print(arg + " ")
	}
	fmt.Println()

	// _ is used to ignore the initialization
	//Args[1:]=> gives a subslice after 1 to the end
}