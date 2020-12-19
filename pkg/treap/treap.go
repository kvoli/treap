package treap

// Austen McClernon
// 834063
// Assignment 1 COMP90077

import (
	"fmt"
	"math/rand"

	"github.com/kvoli/COMP90077_ASS/pkg/elm"
)

// Node represents a treap node data structure
type Node struct {
	Priority    int
	Key         int
	Val         *elm.Element
	Left, Right *Node
}

// OperationEval takes an operation and performs the corresponding operation
func OperationEval(n *Node, o *elm.Operation) *Node {
	switch o.Type {
	case elm.InsertOperation:
		return Insert(n, o.Element)
	case elm.DeleteOperation:
		return Delete(n, o.Element.Key)
	case elm.SearchOperation:
		Search(n, o.Element.Key)
	}
	return n
}

// Insert inserts a new element into the treap and returns the root
func Insert(n *Node, e *elm.Element) *Node {

	// if root == NULL then create new node and return
	if n == nil {
		return newNode(e)
	}

	// check if key is smaller than root
	if e.Key <= n.Key {
		n.Left = Insert(n.Left, e)

		// check if the left subtree has lower priority (minheap)
		if n.Left.Priority > n.Priority {
			n = rRotate(n)
		}
	} else {
		n.Right = Insert(n.Right, e)

		// check if the right subtree has lower priority (minheap)
		if n.Right.Priority > n.Priority {
			n = lRotate(n)
		}
	}

	return n
}

// InOrder returns the in order traversal of the tree
func InOrder(n *Node, depth int) {
	if n != nil {
		InOrder(n.Left, depth+1)
		fmt.Printf("key: %d: | priority: %d | depth: %d", n.Key, n.Priority, depth)
		fmt.Println()
		InOrder(n.Right, depth+1)
	}
}

// MaxDepth calculates the deepest node
func MaxDepth(n *Node) int {
	if n == nil {
		return 0
	}
	l := MaxDepth(n.Left)
	r := MaxDepth(n.Right)
	if l >= r {
		return l + 1
	}

	return r + 1
}

// Delete removes a node from the Treap
func Delete(n *Node, key int) *Node {

	if n == nil {
		return n
	}
	if key < n.Key {
		n.Left = Delete(n.Left, key)
	} else if key > n.Key {
		n.Right = Delete(n.Right, key)
	} else if n.Left == nil {
		n = n.Right
	} else if n.Right == nil {
		n = n.Left
	} else if n.Left.Priority < n.Right.Priority {
		n = lRotate(n)
		n.Left = Delete(n.Left, key)
	} else {
		n = rRotate(n)
		n.Right = Delete(n.Right, key)
	}
	return n
}

// Search does a binary search for the key and returns it if it exists
func Search(n *Node, key int) *Node {

	if n == nil || n.Key == key {
		return n
	}
	if key > n.Key {
		return Search(n.Right, key)
	}

	return Search(n.Left, key)

}

func rRotate(n *Node) *Node {
	x := n.Left
	T2 := x.Right

	x.Right = n
	n.Left = T2

	return x
}

func lRotate(n *Node) *Node {
	y := n.Right
	T2 := y.Left

	y.Left = n
	n.Right = T2

	return y
}

func newNode(e *elm.Element) *Node {
	return &Node{
		Priority: int(rand.Int31n(MaximumPriority)),
		Key:      e.Key,
		Val:      e,
	}
}
