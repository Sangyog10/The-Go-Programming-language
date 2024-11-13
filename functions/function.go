package main

import "fmt"

// func main(){
// 	fmt.Println(add(4,5))
// 	fmt.Println(random())
// }


// func add(x int, y int ) int {
// 	return x+y
// }

// func random()string{
// 	return "sangyog"
// }

// new conditon

// func divide(x , y int) (int, error){
// 	if y==0 {
// 		return 0.0, fmt.Errorf("cannot divide by 0")
// 	}
// 	return x/y, nil
// }

// func main(){
// 	result, err := divide(9,2)
	
// 	if err != nil{
// 		fmt.Println("Error", err)
// 	}else{
// 		fmt.Println("Result is :", result)
// 	}
// }

//return type declared
// func mul(x,y int)( result int){
// 	result = x*y
// 	return 
// }

// func main(){
// 	fmt.Println(mul(4,5))
// }

//Variadic Functions(multiple parameter)

// func add ( numbers ...int)(total int){
// 	total = 0
// 	for _, number := range numbers {
// 		total+=number
// 	} 
// 	return 
// }

// func main(){
// 	fmt.Println(add(3,4,5,5))
// }


// Anonymous Functions and Closures
// func main(){
// 	multiplier := 2
// 	double := func(x int)int{
// 		return x * multiplier
// 	}
// 	//double is closure
// 	fmt.Println("Product is:",double(5))
// }

