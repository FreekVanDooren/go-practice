package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	values1 := make(map[int]bool)
	values2 := make(map[int]bool)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		values1[<-ch1] = true
		values2[<-ch2] = true
	}
	for value := range values1 {
		if !values2[value] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t := tree.New(1)
	fmt.Println(t)
	go Walk(t, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
