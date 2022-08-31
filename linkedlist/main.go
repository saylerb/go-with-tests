package linkedlist

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	var currentNode *SinglyLinkedListNode = llist
	var currentPosition int32

	for currentPosition = 0; currentPosition < position-1; currentPosition++ {
		currentNode = currentNode.next
	}

	nodeToInsert := SinglyLinkedListNode{data: data, next: currentNode.next}

	currentNode.next = &nodeToInsert

	return llist
}
