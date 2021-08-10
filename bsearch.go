package main

import "fmt"

var count int

type Node struct {
	Key   int
	Right *Node
	Left  *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		//Right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{}

	tree.Insert(100)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(600)
	tree.Insert(700)
	tree.Insert(800)

	fmt.Println(tree)
	fmt.Println(tree.Search(300))
	fmt.Println(count)

}
