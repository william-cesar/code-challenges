// Singly-linked lists are already defined with this interface:
//	type ListNode struct {
//	  Value interface{}
//	  Next *ListNode
//	}

func solution(l *ListNode, k int) *ListNode {
	h := &ListNode{}
	h.Next = l
	l = h

	for l.Next != nil {
		if l.Next.Value == k {
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}

	return h.Next
}