package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		k := 5
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Println(i, k)
			m[i] = k
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Printf("%#v\n", m)
}
