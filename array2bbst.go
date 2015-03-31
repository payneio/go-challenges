package main

import (
	"fmt"
)

// convert a sorted array to a balanced binary tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%d\n%s%s", n.Value, n.Left, n.Right)
}

func main() {
	a := []int{1, 2, 4, 6, 8, 12, 14, 15}
	bst := makeBST(a)
	fmt.Println(bst)
}

func makeBST(a []int) *Node {
	if len(a) == 0 {
		return nil
	}
	var n *Node
	if len(a) == 1 {
		n = &Node{Value: a[0]}
	} else {
		mid := len(a) / 2
		n = &Node{Value: a[mid]}
		n.Left = makeBST(a[:mid])
		n.Right = makeBST(a[mid+1:])
	}
	return n
}
