package main

import (
	"fmt"
	"sync"
)

func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2) //we set conter to value=2,
	go func() {
		v = 1
		wg.Done() //end this goroutine, decrease counter value
	}()
	go func() {
		fmt.Println(v) //propably print value=0,sometimes value=1 its bad solution.
		wg.Done()
	}()
	wg.Wait() //waiting until all goroutines ends, until counter=0,if counter is bigger then goroutines count, It will happen error:fatal error: all goroutines are asleep - deadlock!
}
