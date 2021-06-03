package concurrency

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	walkWithoutClose(t, ch)
	close(ch)
}

func walkWithoutClose(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkWithoutClose(t.Left, ch)
	ch <- t.Value
	walkWithoutClose(t.Right, ch)
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int),make(chan int)
	go walk(t1, ch1)
	go walk(t2, ch2)

	eq := true
	for i := range ch1 {
		v, ok := <- ch2
		eq = eq && ok && (i == v)
	}
	_, ok := <- ch2
	eq = eq && !ok

	return eq
}

func SameTree() {
	tree1 := tree.New(7)
	fmt.Println(tree1)

	tree2 := tree.New(7)
	fmt.Println(tree2)
	fmt.Println("same ", Same(tree1, tree2))
}
