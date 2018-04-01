package main

import "fmt"

func main() {
	tree := New()
	fmt.Println("Insert 15 1 2")
	tree.Insert(15)
	tree.Insert(1)
	tree.Insert(2)
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Show: ")
	tree.Show()
	fmt.Println("Insert 3 1 3")
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(3)
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Show: ")
	tree.Show()
	fmt.Println("Search 4: ", tree.Search(4))
	fmt.Println("Search 3: ", tree.Search(3))
	fmt.Println("Insert 14 17 31")
	tree.Insert(14)
	tree.Insert(17)
	tree.Insert(31)
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Show: ")
	tree.Show()
	fmt.Println("Min element: ")
	tree.FindMin()
	fmt.Println("Max element: ")
	tree.FindMax()
	fmt.Println("Delete 31: ", tree.Delete(31))
	fmt.Println("Delete 5: ", tree.Delete(5))
	fmt.Println("Delete 5: ", tree.Delete(1))
	fmt.Println("Size: ", tree.Size())
	fmt.Println("Show: ")
	tree.Show()
	fmt.Println("Min element: ")
	tree.FindMin()
	fmt.Println("Max element: ")
	tree.FindMax()
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
* + Delete(k): удаление элемента k.
* + Search(k): поиск значения элемента k в структуре, есть он или нет.
* + FindMax(): поиск максимального значения.
* + FindMin(): поиск минимального значения.
* + Show & Size(): печать дерева и размер.
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

// Show the tree (Print the tree in-order)
func (tree *Bst) Show() {
	printNode(tree.root)
}

// Print the tree in-order
// Traverse the left sub-tree, root, right sub-tree
func printNode(root *Node) {
	if root != nil {
		printNode(root.left)
		fmt.Println(root.key)
		printNode(root.right)
	}
}

// FindMin - print min element tree
func (tree *Bst) FindMin() {
	fmt.Println(minValue(tree.root))
}

func minValue(root *Node) int {
	if root != nil {
		if root.left == nil {
			return root.key
		}
		return minValue(root.left)
	}
	return root.key
}

// FindMax - print max element tree
func (tree *Bst) FindMax() {
	fmt.Println(maxValue(tree.root))
}

func maxValue(root *Node) int {
	if root != nil {
		if root.right == nil {
			return root.key
		}
		return maxValue(root.right)
	}
	return root.key
}

// Delete element tree
func (tree *Bst) Delete(value int) bool {
	if !tree.Search(value) || tree.root == nil {
		return false
	}

	if tree.root.key == value {
		tempRoot := &Node{0, nil, nil}
		tempRoot.left = tree.root
		r := del(tree.root, tempRoot, value)
		tree.root = tempRoot.left
		return r
	}
	return del(tree.root.left, tree.root, value) || del(tree.root.right, tree.root, value)
}

func del(root *Node, parent *Node, value int) bool {
	switch {
	case root.key == value:
		if root.left != nil && root.right != nil {
			root.key = minValue(root.right)
			return del(root.right, root, root.key)
		}
		link(parent, root)
		return true
	case root.key > value:
		if root.left == nil {
			return false
		}
		return del(root.left, root, value)
	case root.key < value:
		if root.right == nil {
			return false
		}
		return del(root.right, root, value)
	}
	return false
}

func link(parent *Node, root *Node) {
	if parent.left == root {
		if root.left != nil {
			parent.left = root.left
		} else {
			parent.left = root.right
		}
	} else if parent.right == root {
		if root.left != nil {
			parent.right = root.left
		} else {
			parent.right = root.right
		}
	}
}
