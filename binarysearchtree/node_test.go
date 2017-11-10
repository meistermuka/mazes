package binarytree

import (
	"testing"
)

var bst ItemBinarySearchTree

func fillTree(bst *ItemBinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(1, "1")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(3, "3")
}

func TestInsert(t *testing.T) {
	fillTree(&bst)
	bst.String()

	bst.Insert(11, "11")
	bst.String()
}
