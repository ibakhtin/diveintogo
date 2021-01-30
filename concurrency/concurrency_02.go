// +build OMIT

package main

import "fmt"

func logger(channel <-chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func main() {
	logChannel := make(chan string)
	defer close(logChannel)

	go logger(logChannel)

	for i := 0; i < 100; i++ {
		logChannel<- fmt.Sprint(i)
	}
}
