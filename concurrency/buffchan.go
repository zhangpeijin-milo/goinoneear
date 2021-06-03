package concurrency

import "fmt"

func Full()  {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)

	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
