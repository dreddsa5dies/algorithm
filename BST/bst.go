package main

import "fmt"

func main() {
	tree := New()
	fmt.Println("Insert 15 1 2")
	tree.Insert(15)
	tree.Insert(1)
	tree.Insert(2)
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Insert 3 1 3")
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(3)
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Search 4: ", tree.Search(4))
	fmt.Println("Search 3: ", tree.Search(3))
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
* + Insert(k): вставка элемента k в дерево.
* Delete(k): удаление элемента k.
* + Search(k): поиск значения элемента k в структуре, есть он или нет.
* FindMax(): поиск максимального значения.
* FindMin(): поиск минимального значения.
*/

// New - Construtor BST
func New() *Bst {
	return &Bst{nil, 0}
}

// Insert elements in tree
func (tree *Bst) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{value, nil, nil}
	}
	tree.size++
	tree.root.insert(&Node{value, nil, nil})
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

// Size - return size tree
func (tree *Bst) Size() int {
	return tree.size
}

// Search element on tree
func (tree *Bst) Search(value int) bool {
	tree.size--
	return searchElement(tree.root, value)
}

// search element
func searchElement(root *Node, value int) bool {
	if root != nil {
		if value == root.key {
			return true
		} else if value > root.key {
			return searchElement(root.right, value)
		} else {
			return searchElement(root.left, value)
		}
	}
	return false
}

// Print the tree in-order
// Traverse the left sub-tree, root, right sub-tree
func showInOrder(root *Node) {
	if root != nil {
		showInOrder(root.left)
		fmt.Println(root.key)
		showInOrder(root.right)
	}
}

// Print the tree pre-order
// Traverse the root, left sub-tree, right sub-tree
func showPreOrder(root *Node) {
	if root != nil {
		fmt.Println(root.key)
		showInOrder(root.left)
		showInOrder(root.right)
	}
}

// Print the tree post-order
// Traverse left sub-tree, right sub-tree, root
func showPostOrder(root *Node) {
	if root != nil {
		fmt.Println(root.key)
		showInOrder(root.left)
		showInOrder(root.right)
	}
}
