// Singly-linked lists are already defined with this interface:
//	type ListNode struct {
//	  Value interface{}
//	  Next *ListNode
//	}

func solution(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Value.(int) <= l2.Value.(int) {
		l1.Next = solution(l1.Next, l2)
		return l1
	}

	l2.Next = solution(l1, l2.Next)
	return l2
}