package concurrency

import (
	"fmt"
)

func sum(s []int, c chan int)  {
	sum := 0
	for _,v := range s {
		sum += v
	}

	c <- sum
}

func Sum(s []int){
	l := len(s)

	c := make(chan int)
	go sum(s[:l/2], c)
	go sum(s[l/2:], c)

	x := <-c
	y := <-c

	fmt.Println(x, y)
}