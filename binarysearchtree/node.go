package binarytree

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item is the type of the BST
type Item generic.Type

// Node is a simple node that composes a tree
type Node struct {
	key   int
	value Item
	left  *Node
	right *Node
}

// ItemBinarySearchTree is the BST of Items
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

// Insert the item inside the tree
func (bst *ItemBinarySearchTree) Insert(key int, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// internal function doing the actual insert
func insertNode(node *Node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse visits all the nodes in-order traversing
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	inOrderTraverse(bst.root, f)
}

// internal recursive function for InOrderTraverse
func inOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse visits all nodes in pre-order traversing
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	preOrderTraverse(bst.root, f)
}

// internal resursive function for PreOrderTraverse
func preOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse visits all nodes in post-order traversing
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

// Min returns the Item with the smallest value stored in the tree
func (bst *ItemBinarySearchTree) Min() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max returns the Item with the biggest value stored in the tree
func (bst *ItemBinarySearchTree) Max() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.right == nil {
			return &n.value
		}

		n = n.right
	}
}

// Search returns true if an item exists in the tree
func (bst *ItemBinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

// internal recursive search function
func search(n *Node, key int) bool {

	if n == nil {
		return false
	}

	if key < n.key {
		return search(n.left, key)
	}

	if key > n.key {
		return search(n.right, key)
	}

	return true
}

// Remove removes an Item from the tree based on a key
func (bst *ItemBinarySearchTree) Remove(key int) *Node {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return remove(bst.root, key)
}

func remove(n *Node, key int) *Node {

	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = remove(n.left, key)
		return n
	}

	if key > n.key {
		n.right = remove(n.right, key)
		return n
	}
	// key == n.key
	if n.left == nil && n.right == nil {
		n = nil
		return nil
	}

	if n.left == nil {
		n = n.right
		return n
	}

	if n.right == nil {
		n = n.left
		return n
	}

	leftmostrightside := n.right
	for {
		// find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	n.key, n.value = leftmostrightside.key, leftmostrightside.value
	n.right = remove(n.right, n.key)

	return n
}

//String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("---------------------------------")
	stringify(bst.root, 0)
	fmt.Println("---------------------------------")

}

func stringify(n *Node, level int) {

	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "    "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}
