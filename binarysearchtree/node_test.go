package binarytree

import (
	"fmt"
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

func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestInOrderTraverse(t *testing.T) {
	var result []string
	bst.InOrderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})

	if !isSameSlice(result, []string{"1", "2", "3", "4", "6", "8", "10", "11"}) {
		t.Errorf("Traversal order incorrect, got %v", result)
	}
}

func TestMin(t *testing.T) {
	if fmt.Sprintf("%s", *bst.Min()) != "1" {
		t.Errorf("MIN should be 1")
	}
}
