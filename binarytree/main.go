package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func preOrderMut(root *Node) string {
	var data []int
	data = traverseMut(&data, root)

	return printData(data)
}

func preOrderCopying(root *Node) string {
	var data []int
	data = traverseCopying(data, root)

	return printData(data)
}

func printData(data []int) string {
	var result string
	for index, number := range data {
		if index != 0 {
			result += " "
		}
		result += fmt.Sprint(number)
	}
	return result
}

func traverseMut(state *[]int, node *Node) []int {
	*state = append(*state, node.data)
	if node.left != nil {
		traverseMut(state, node.left)
	}

	if node.right != nil {
		traverseMut(state, node.right)
	}

	return *state
}

func traverseCopying(state []int, node *Node) []int {
	result := append(state, node.data)
	if node.left != nil {
		result = traverseCopying(result, node.left)
	}

	if node.right != nil {
		result = traverseCopying(result, node.right)
	}

	return result
}
