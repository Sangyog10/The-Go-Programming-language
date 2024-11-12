package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	var s, sep string
	for i :=1;  i< len(os.Args); i++{
		s += sep + os.Args[i]
		sep=" "
	}
	fmt.Println(s)
	alternative()
}


//alternative approach: string is immutable, so it will take longer time to add on string , so join func is faster(it allocates for the final o/p)

func alternative(){
	fmt.Println(strings.Join(os.Args[1:], " "))
}

