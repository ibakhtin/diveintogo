package main

import "fmt"

func greeter(in chan string, out chan string) {
	value := <- in
	fmt.Println(value)
	out <- "Hello, " + value + "!"
}

func main() {
	in := make(chan string)
	out := make(chan string)

	go greeter(in, out)

	in <- "Igor"

	fmt.Println(<- out)
}
