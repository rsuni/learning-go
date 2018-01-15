package main

import (
	"fmt"
	"sync"
)

var count int
var m sync.Mutex

type Writer interface {
	Write(message string)
}

type Logger1 struct {
	Writer
}
type Logger2 struct {
	Writer
}

func (logger Logger1) Write(message string) {

	fmt.Printf("Logger1:%s \n", message)
}
func (logger Logger2) Write(message string) {

	fmt.Printf("Logger2:%s \n", message)
}

func createValue(prefix string, ch chan string) {
	counter := 10

	m.Lock()
	count = count + counter
	m.Unlock()

	for i := 0; i < counter; i++ {
		value := fmt.Sprintf("%s:%d", prefix, i)
		ch <- value
	}

}
func printLogs(ch chan string, done chan bool) {

	var loggers []Writer
	l1 := Logger1{}
	l2 := Logger2{}
	loggers = append(loggers, l1)
	loggers = append(loggers, l2)

	for {
		value := <-ch
		m.Lock()
		count--
		m.Unlock()

		for _, logger := range loggers {
			logger.Write(value)
		}
		if count <= 0 {
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan string)
	done := make(chan bool, 1)

	go createValue("a", ch)
	go createValue("b", ch)
	go createValue("c", ch)
	go printLogs(ch, done)
	<-done
}
