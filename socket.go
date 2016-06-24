package main

import (
	"fmt"
	"time"
)

var (
	readFile = "RECEIVE %-12v: %v\n"
	writeF   = "SEND    %-12v: %v\n"
)

// recieve ...
func recieve(c <-chan string, prefix string) {
	for d := range c {
		fmt.Printf(readFile, prefix, d)
	}
}

// send ...
func send(c chan<- string, data, prefix string) {
	for _, s := range data {
		fmt.Printf(writeF, prefix, string(s))
		c <- string(s)
	}
}

func main() {
	unBChan := make(chan string)
	bChan := make(chan string, 10)
	defer close(bChan)
	defer close(unBChan)

	fmt.Println("Start working")
	bChan <- "Start bChan"
	fmt.Printf(readFile, "First bChan", <-bChan)

	go recieve(bChan, "First")

	send(bChan, "some text", "First")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n------------------------")

	go recieve(unBChan, "Second")

	send(unBChan, "second text", "Second")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("End working")
}
