package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hello")
	go say("hi")
	go hello()
	time.Sleep(time.Second * 1)

}
func hello() {
	fmt.Println("hello")
}
func say(s string) {
	for i := 0; i > 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 500)
	}

}
