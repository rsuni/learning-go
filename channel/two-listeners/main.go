package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	ch := make(chan string)

	go func() {

		for v := range ch {
			fmt.Println("1:", v)
			wg.Done()
		}

	}()
	go func() {
		for v := range ch {
			fmt.Println("2:", v)
			wg.Done()
		}
	}()
	go func() {
		for v := range ch {
			fmt.Println("3:", v)
			wg.Done()
		}
	}()
	go func() {
		for v := range ch {
			fmt.Println("4:", v)
			wg.Done()
		}
	}()
	go func() {
		for v := range ch {
			fmt.Println("5:", v)
			wg.Done()
		}
	}()

	ch <- "hello"
	ch <- "world"
	ch <- "how"
	ch <- "are you"
	wg.Wait()

	return
}
