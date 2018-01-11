package main

import (
	"fmt"
	"sync"
)

type Processor struct {
	mu  sync.Mutex
	sum int
}

var wg sync.WaitGroup
var mu sync.Mutex

func process(n string, p *Processor) {

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 100000; i++ {
			p.mu.Lock()
			p.sum = p.sum + 1 //only one process has access.

			p.mu.Unlock()

		}

		fmt.Printf("From %s:sum=%d \n", n, p.sum)

	}()
}

func main() {
	p1 := Processor{}
	p1.sum = 0

	p2 := Processor{}
	p2.sum = 0

	processes := []string{"A", "B", "C", "D", "E"}
	for _, p := range processes {
		process(p, &p1)
	}

	processes = []string{"F", "G", "H", "I", "J"}
	for _, p := range processes {
		process(p, &p2)
	}
	wg.Wait()
	fmt.Println("Final Sum:", p1.sum)
	fmt.Println("Final Sum2:", p2.sum)
}
