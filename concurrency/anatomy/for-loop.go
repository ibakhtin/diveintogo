// +build OMIT

package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)
	go squares(c)
	for {
		val, ok := <-c
		if ok {
			fmt.Println(val, ok)
		} else {
			fmt.Println(val, ok, "<-- loop broke!")
			break
		}
	}
	fmt.Println("main() stopped")
}
