package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	// go greet("How are you?", dones[1])
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	// go slowGreet("How ... are ... you ...?", dones[2])
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	// go greet("I hope you're liking the course!", dones[3])
	go greet("I hope you're liking the course!", done)
	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
	}
}
