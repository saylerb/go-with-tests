package main

import "testing"

func TestBinaryTree(t *testing.T) {
	// 1
	//  \
	//   2
	//    \
	//     5
	//    /  \
	//   3    6
	//    \
	//     4
	node6 := Node{data: 6, left: nil, right: nil}
	node4 := Node{data: 4, left: nil, right: nil}
	node3 := Node{data: 3, left: nil, right: &node4}
	node5 := Node{data: 5, left: &node3, right: &node6}
	node2 := Node{data: 2, left: nil, right: &node5}
	root := Node{data: 1, left: nil, right: &node2}

	t.Run("preorder traversal mutating state", func(t *testing.T) {
		got := preOrderMut(&root)
		want := "1 2 5 3 4 6"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("preorder traversal with immutable state", func(t *testing.T) {
		got := preOrderCopying(&root)
		want := "1 2 5 3 4 6"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
