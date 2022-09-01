package linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertIntoLinkedList(t *testing.T) {
	t.Run("insert 1 at position 2 of list of length 3", func(t *testing.T) {
		tail := SinglyLinkedListNode{data: 7, next: nil}
		middle := SinglyLinkedListNode{data: 13, next: &tail}
		head := SinglyLinkedListNode{data: 16, next: &middle}

		var result *SinglyLinkedListNode

		result = insertNodeAtPosition(&head, 1, 2)

		got := convertLinkedListToString(result, " ", t)
		want := "16 13 1 7"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("insert 20 at position 3 of list of length 5", func(t *testing.T) {
		tail := SinglyLinkedListNode{data: 4, next: nil}
		four := SinglyLinkedListNode{data: 10, next: &tail}
		three := SinglyLinkedListNode{data: 19, next: &four}
		two := SinglyLinkedListNode{data: 9, next: &three}
		head := SinglyLinkedListNode{data: 11, next: &two}

		var result *SinglyLinkedListNode

		result = insertNodeAtPosition(&head, 20, 3)

		got := convertLinkedListToString(result, " ", t)
		want := "11 9 19 20 10 4"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("insert 7 at position 5 of list of length 6", func(t *testing.T) {
		tail := SinglyLinkedListNode{data: 6, next: nil}
		five := SinglyLinkedListNode{data: 5, next: &tail}
		four := SinglyLinkedListNode{data: 4, next: &five}
		three := SinglyLinkedListNode{data: 3, next: &four}
		two := SinglyLinkedListNode{data: 2, next: &three}
		head := SinglyLinkedListNode{data: 1, next: &two}

		var result *SinglyLinkedListNode

		result = insertNodeAtPosition(&head, 7, 5)

		got := convertLinkedListToString(result, " ", t)
		want := "1 2 3 4 5 7 6"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestDeleteNodeFromLinkedList(t *testing.T) {
	t.Run("remove node from postion 1 from list of length 4", func(t *testing.T) {
		head := SinglyLinkedListNode{
			data: 11,
			next: &SinglyLinkedListNode{
				data: 9,
				next: &SinglyLinkedListNode{
					data: 2,
					next: &SinglyLinkedListNode{
						data: 9,
						next: nil,
					},
				},
			},
		}

		result := deleteNode(&head, 1)

		got := convertLinkedListToString(result, " ", t)
		want := "11 2 9"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("remove the head from a linked list ", func(t *testing.T) {
		head := SinglyLinkedListNode{
			data: 5,
			next: &SinglyLinkedListNode{
				data: 4,
				next: &SinglyLinkedListNode{
					data: 3,
					next: nil,
				},
			},
		}

		result := deleteNode(&head, 0)

		got := convertLinkedListToString(result, " ", t)
		want := "4 3"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func convertLinkedListToString(node *SinglyLinkedListNode, sep string, t *testing.T) string {
	t.Helper()
	var result string

	for node != nil {
		result += fmt.Sprint(node.data)
		node = node.next

		if node != nil {
			result += sep
		}
	}

	return result
}
