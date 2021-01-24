// +build OMIT

package main

import "fmt"

func greet(c chan string) {
	fmt.Println(<-c) // for John
	fmt.Println(<-c) // for Mike
}

func main() {
	fmt.Println("main() started")
	c := make(chan string, 1)
	go greet(c)
	c <- "John"
	close(c) // closing channel
	c <- "Mike"
	fmt.Println("main() stopped")
}
