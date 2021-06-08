package main

import (
	"fmt"
	"time"
)

func main() {
	go producer(1)
	go producer(2)
	time.Sleep(time.Millisecond * 1)
}
func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf(" Producer # %v - message # %v\n", id, i)
	}
}
