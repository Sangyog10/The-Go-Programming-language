package main

import "fmt"

func main() {
	sum := add(4, 5)
	concat := add("sang", "yog")
	fmt.Print(sum, concat)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
