package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)
		go fmt.Println(string(data))
		time.Sleep(time.Millisecond * 1)
	}
}
