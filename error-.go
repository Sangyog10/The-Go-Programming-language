package main

import (
	"fmt"
	"os"
	"errors"
)

func check (a, b int) error{
	if a==0 && b==0{
		return errors.New("this is a custom error message")
	}
	return nil
}

func formattedError(a, b int) error{
	if a==b && b==0 {
		return fmt.Errorf("a = %d and b = %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func main(){
	err := check(0,10)
	if err == nil{
		fmt.Println("check() executed normally!")
	}else{
		fmt.Println(err)
	}

	err = check(0,0)
	if err.Error() == "this is a custom error message"{
		fmt.Println("Custom error detected!")
	}

	err = formattedError(0,0)
	if err != nil{
		fmt.Println(err)
	}

}