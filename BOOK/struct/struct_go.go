package main

import "fmt"

type Location struct{
	city string
	zipcode int
}

type Person struct{
	Name string
	age int
	Location
}


func main(){
	p := Person{
		Name:"sangyog", 
		age: 22,
		Location: Location{city:"sindhuli", zipcode:44900},
	}
	fmt.Println("Person:",p)
	p.age = 23
	fmt.Println("Person's age:",p.zipcode)
}