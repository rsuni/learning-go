package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sum = 0

func process(n string) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10000; i++ {
			sum = sum + 1
		}

		fmt.Println("From "+n+":", sum)
	}()
}

func main() {
	processes := []string{"A", "B", "C", "D", "E"}
	for _, p := range processes {
		process(p)
	}

	wg.Wait()
	fmt.Println("Final Sum:", sum) //Finally sum value can be lesser then 50000 because sometime 2 processes write in same time
}
