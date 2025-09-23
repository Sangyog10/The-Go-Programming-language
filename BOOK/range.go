package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}



func main() {

	for i, v := range pow{
		fmt.Println(i, v)
	}
}

//range returns the index and value at that index(i,v)


func main() {
	pow := make([]int , 10)
	for i := range pow{
		fmt.Println(i) //returns index
	}
	for _,value := range pow{
		fmt.Println(value) //returns value at each index

	}
}
