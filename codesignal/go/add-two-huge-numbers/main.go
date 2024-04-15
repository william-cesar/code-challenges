// Singly-linked lists are already defined with this interface:
//	type ListNode struct {
//	  Value interface{}
//	  Next *ListNode
//	}

func solution(a *ListNode, b *ListNode) *ListNode {
	aSlice := make([]int, 0)
	for a != nil {
		aSlice = append(aSlice, a.Value.(int))
		a = a.Next
	}
	bSlice := make([]int, 0)
	for b != nil {
		bSlice = append(bSlice, b.Value.(int))
		b = b.Next
	}
	if len(aSlice) < len(bSlice) {
		aSlice, bSlice = bSlice, aSlice
	}
	var prev *ListNode
	c := 0
	for i, j := len(aSlice)-1, len(bSlice)-1; i >= 0; i, j = i-1, j-1 {
		sum := aSlice[i]
		if j >= 0 {
			sum += bSlice[j]
		}
		sum += c
		val := sum % 10000
		c = sum / 10000
		node := ListNode{val, prev}
		prev = &node
	}
	if c > 0 {
		node := ListNode{c, prev}
		prev = &node
	}
	return prev
}