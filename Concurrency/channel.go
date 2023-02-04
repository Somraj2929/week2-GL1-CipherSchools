package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) { //sender channel
		ch <- "2"
		fmt.Println(1)
	}(channel)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
}

func main1() {
	channel := make(chan string, 1)
	go func(ch chan string) { //bidirectional channel
		mess := <-ch
		ch <- "hey from anonymous"
		fmt.Println(mess)
		fmt.Println(1)
	}(channel)
	message := "Hello from MAIN FUNCTION"
	channel <- message
	time.Sleep(time.Second * 5)
	message = <-channel
	fmt.Println(message)
}
