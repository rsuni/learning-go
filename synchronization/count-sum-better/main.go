package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var sum = 0
var sum2 = 0

func process(n string) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10000; i++ {
			mu.Lock()
			sum = sum + 1 //only one process has access.
			sum2 = sum2 + 1
			mu.Unlock()
		}

		//fmt.Println("From "+n+":", sum)
		mu.Lock() //without this lock, in next print is sum2 bigger then sum.
		fmt.Printf("From %s:sum=%d,sum2:%d \n", n, sum, sum2)
		mu.Unlock()
	}()
}

func main() {
	processes := []string{"A", "B", "C", "D", "E"}
	for _, p := range processes {
		process(p)
	}

	wg.Wait()
	fmt.Println("Final Sum:", sum)
	fmt.Println("Final Sum2:", sum2)
}
