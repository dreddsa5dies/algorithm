package main

import (
	"fmt"

	"github.com/dreddsa5dies/algorithm/util"
)

func main() {
	list := util.RandomInt() // срез int
	fmt.Printf("List:\t%v\n", list)

	tree := New()
	fmt.Println(tree)
	tree.Insert(10)
	tree.Insert(0)
	fmt.Println(tree.size)
	fmt.Println(tree.root)
	fmt.Println(tree)
}

// Node is a representation of a single node in tree. (recursive ADT)
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Bst is a struct of a nodes in tree.
type Bst struct {
	root *Node
	size int
}

/*
Binary Search Tree ADT Operations
* Insert(k): вставка элемента k в дерево.
* Delete(k): удаление элемента k.
* Search(k): поиск значения элемента k в структуре, есть он или нет.
* FindMax(): поиск максимального значения.
* FindMin(): поиск минимального значения.
*/

// New - Construtor BST
func New() *Bst {
	return &Bst{nil, 0}
}

// insert is a recursive method for node insertion
func (root *Node) insert(newNode *Node) {

	//if data exists, skip
	if root.key == newNode.key {
		return
	}

	// to right-subtree
	if root.key < newNode.key {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

// Insert elements in tree
func (tree *Bst) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{value, nil, nil}
	}
	tree.size++
	tree.root.insert(&Node{value, nil, nil})
}
