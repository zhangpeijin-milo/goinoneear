package concurrency

import (
	"fmt"
	"time"
)

func world() {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("world")
}

func hello() {
	fmt.Println("hello")
}

func HelloWorld() {
	go world()
	hello()
}