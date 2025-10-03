package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// doubleResult := transformNumbers(&numbers, double)
	// tripleResult := transformNumbers(&numbers, triple)
	// fmt.Println(doubleResult)
	// fmt.Println(tripleResult)
	// getTransformerFunc(&numbers)
	// fmt.Println(AnanymousFunc())

	// doubledValue := createTransformer(2)
	// tripledValue := createTransformer(3)
	// fmt.Println(doubledValue, tripledValue)

	//variadic function
	sumup(1, 2, 3, 4, 5)
	// sumup(numbers) //it won't work for variadic func
	sumup(numbers...)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNums := []int{}
	for _, value := range *numbers {
		dNums = append(dNums, transform(value))
	}
	return dNums
}

func getTransformerFunc(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}

// Ananymous function: Can be used when we do not need to use it elsewhere and can instantly define it and use it

func AnanymousFunc() []int {
	numbers := []int{1, 2, 3, 4, 5}
	return transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
}

// Closure: Use a specific aspect of ananynomous func
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// variadic function: func with any number of parameter

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
