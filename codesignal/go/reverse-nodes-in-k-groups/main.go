// Singly-linked lists are already defined with this interface:
//	type ListNode struct {
//	  Value interface{}
//	  Next *ListNode
//	}

func solution(l *ListNode, k int) *ListNode {
	var count int

	head := l
	current := head
	sliceHead := head

	for current != nil {
		if count == k-1 && k != 1 {
			next := current.Next
			current.Next = nil
			if head == l {
				head = reverse(sliceHead)
				sliceHead.Next = next
			} else {
				nextSliceHead := sliceHead.Next
				sliceHead.Next = reverse(sliceHead.Next)
				sliceHead = nextSliceHead
				sliceHead.Next = next
			}
			count = 0
			current = next
		} else {
			count++
			current = current.Next
		}
	}

	return head
}

func reverse(l *ListNode) *ListNode {
	var head *ListNode

	for l != nil {
		next := l.Next
		l.Next = head
		head = l
		l = next
	}

	return head
}
