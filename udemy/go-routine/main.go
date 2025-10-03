package main

import (
	"fmt"
	"time"
)

/*Using channel individually
func main() {
	dones := make([]chan bool, 3)
	// done := make(chan bool)

	dones[0] = make(chan bool)
	go fastGreet("This is is first operation", dones[0])
	dones[1] = make(chan bool)
	go slowPrint("This is secodn operaton and  i will take time", dones[1])
	dones[2] = make(chan bool)
	go greet("This is third operation ", dones[2])
	// <-done
	// <-done
	// <-done

	for _, done := range dones {
		<-done
	}
}
*/

func main() {
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go fastGreet("This is is first operation", done)
	// dones[1] = make(chan bool)
	go slowPrint("This is secodn operaton and  i will take time", done)
	// dones[2] = make(chan bool)
	go greet("This is third operation ", done)

	for donechain := range done {
		fmt.Println(donechain)
	}

}

func slowPrint(str string, doneChan chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println(str)
	doneChan <- true
}
func greet(str string, doneChan chan bool) {
	// time.Sleep(1 * time.Second)
	fmt.Println(str)
	doneChan <- true
}
func fastGreet(str string, doneChan chan bool) {
	// time.Sleep(1 * time.Second)
	fmt.Println(str)
	doneChan <- true
}
