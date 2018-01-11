package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 1) //we can even wait for a while, second goiroutine waiting for values from channel
		v = 1
		ch <- 11
		wg.Done()
	}()

	go func() {
		v2 := <-ch
		fmt.Println(v)
		fmt.Println(v2)
		wg.Done()
	}()
	wg.Wait()

}
