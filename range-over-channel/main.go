package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int) //its differ where is channel definied, here: only 1 circle
	count := 0
	for count < 3 {
		ch := make(chan int) //here: every circle will be run
		go fibonacciProducer(ch, 4)
		idx := 0
		for num := range ch {
			fmt.Printf("F(%d): \t%d\n", idx, num)
			idx++
		}

		count++
		// if count == 3 {
		// 	close(ch)
		// }
	}
}

func fibonacciProducer(ch chan int, count int) {
	fmt.Println("We waiting for a while, before we generate values to the channel")
	time.Sleep(time.Second * 1) //we can even wait for a while, main procedure waiting for values from channel
	n2, n1 := 0, 1
	for count >= 0 {
		ch <- n2
		count--
		n2, n1 = n1, n2+n1
	}
	close(ch) //without closing channel we get error on line 9: "fatal error: all goroutines are asleep - deadlock!"
}
