package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("Excuting Goroutine")
	time.Sleep(1 * time.Second)
	c <- "Hello World"
	fmt.Println("Finished Executing Goroutine")
}

func main() {
	// set capacity of the channel as 2
	vals := make(chan string, 2)
	defer close(vals)

	go SendValue(vals)
	go SendValue(vals)

	val := <-vals
	fmt.Println(val)

	time.Sleep(1 * time.Second)
}
