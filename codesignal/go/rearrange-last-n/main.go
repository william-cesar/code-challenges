// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }

func solution(l *ListNode, n int) *ListNode {
	if n == 0 {
		return l
	}

	var count int
	pointer1 := l
	pointer2 := l

	for pointer1.Next != nil {
		pointer1 = pointer1.Next
		count++
		if count >= n+1 {
			pointer2 = pointer2.Next
		}
	}

	if pointer2 == l && count < n {
		return l
	}

	head := pointer2.Next
	pointer2.Next = nil
	pointer1.Next = l

	return head
}