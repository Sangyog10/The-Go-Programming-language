package main
import (
	"fmt"	
)

func main(){
	// aString := "Hello world! €"
	// fmt.Println(aString)
	// fmt.Println("First byte in string is ", string(aString[0])) //gives the output in hexadecimal(0 to 255)
	// for i:=0; i<len(aString); i++{
	// 	fmt.Println(aString[i])
	// }

r := '€'
fmt.Println("As an int32 value:", r)
fmt.Printf("As a string: %s and as a character: %c\n", r, r)

}
