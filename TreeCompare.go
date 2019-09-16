package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	"sort"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	ch <- t.Value
	if t.Left != nil{
		go Walk(t.Left, ch)
	}
	if t.Right != nil{
		go Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	t1_ch := make(chan int)
	t2_ch := make(chan int)
	
	go Walk(t1, t1_ch)
	go Walk(t2, t2_ch)
	
	t1_val := make([]int, 10)
	t2_val := make([]int, 10)
	
	for i := 0; i < 10; i++{
		t1_val[i] = <- t1_ch
		t2_val[i] = <- t2_ch
	}
	
	ans := true
	sort.Slice(t1_val, func(i, j int) bool { return t1_val[i] <= t1_val[j]; })
	sort.Slice(t2_val, func(i, j int) bool { return t2_val[i] <= t2_val[j]; })
	for i := range t1_val{
		if t1_val[i] != t2_val[i]{
			ans = false
			break
		}
	}
	
	return ans
}

func main() {
	fmt.Println( Same(tree.New(3), tree.New(1)) )
}
