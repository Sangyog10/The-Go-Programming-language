package main

import "fmt"

func main() {
	// arrayList()
	// slicesInGo()
	// dynamicArrayFromList()
	// maps()
	createSliceWithMakeFunc()
	createMapsWithMakefunc()

}

/*
MEMORY ALLOCATION:
When we make a slice or maps wihout fixed length, it will create an empty map/slice. When we add new, it will
create a new array/slice of new length.
So to fix this,we can use make function to preallocate the space for map/slice. It will work on the same location
for the size of it and will only reallocate new if size of items  exceeds than we have created
*/

func createSliceWithMakeFunc() {
	userNames := make([]string, 2, 5) //create an array of capacity 5 with 2 empty slots, append will work from 3rd slot
	userNames[0] = "sangyog,"
	userNames = append(userNames, ",Max,", "Hari,") //it will append the element in 2nd and 3rd index as 1st and 2nd index is alerady occupied with make func
	fmt.Println(userNames)
}

func createMapsWithMakefunc() {
	websites := make(map[int]string, 5)
	websites[1] = "Hey"
	websites[2] = "hi"
	websites[3] = "HHeey"
	fmt.Print(websites)

}

// map will map the key value pair , it always dynamic
func maps() {
	websites := map[int]string{
		1: "google",
		2: "aws",
	}
	websites[3] = "Linked In"
	fmt.Println(websites[1])

	delete(websites, 1)
	fmt.Println(websites)
}

func dynamicArrayFromList() {
	prices := []float64{10, 20}
	// fmt.Println(len(prices), cap((prices)))
	updatedPrices := append(prices, 10.34) //since slice(array) has fixed lenght, when we try to increse its size, it will create new array and return it behind the scene
	fmt.Println(prices, updatedPrices)

	//checking the address of all
	addrPrice := &prices[0]
	addrUpdate := &updatedPrices[0]
	fmt.Println(addrUpdate, addrPrice)

	//appending the new slice to old slice
	discountedPrice := []float64{1, 2, 4}
	prices = append(updatedPrices, discountedPrice...) // ...will insert each element separately one by one as append needs it
}

// /if we modify the element in slice , it will also modify in the oriignal array
func slicesInGo() {
	prices := [4]float64{10.254, 23, 455.44, 23}
	featuredPrices := prices[1:3]
	featuredPrices[1] = 100
	// fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices))
	fmt.Println(cap(featuredPrices)) //cap is the capacity of the slice, it will omit the values from left as we cannot add element in the left

}

// array or lists
func arrayList() {
	var products [4]string = [4]string{"A book"}
	fmt.Println((products))
	prices := [4]float64{10.254, 23, 455.44, 23}
	fmt.Println((prices))
}
