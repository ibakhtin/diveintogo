// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	go greet(c)
	c <- "John"
	fmt.Println("main() stopped")
}
