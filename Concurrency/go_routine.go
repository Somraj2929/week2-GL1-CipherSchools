package main

import (
	"fmt"
	"sync"
	//"time"
)

//var wait sync.WaitGroup

func main() {
	var wait sync.WaitGroup
	counter := 200
	wait.Add(counter)
	wait.Add(1)
	//go hello(wait)
	//go hello(&wait)
	//go hello()
	//time.Sleep(time.Second*5)
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
	defer wait.Wait()
	//go hello(&wait)
}

func hello(wait *sync.WaitGroup, counter int) {
	fmt.Println(counter) //go routine are not executed in a particular order
	// go func() {
	// 	fmt.Println(10)
	// 	fmt.Println(13)
	// 	fmt.Println(19)
	// }()
	wait.Done()
}
