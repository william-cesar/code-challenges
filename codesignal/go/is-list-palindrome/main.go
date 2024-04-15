// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }

func tolist(l *ListNode) (list []interface{}) {
	for ; l != nil; l = l.Next {
		list = append(list, l.Value)
	}
	return list
}

func solution(l *ListNode) bool {
	list := tolist(l)
	lenlist := len(list)

	for i := 0; i < lenlist/2; i++ {
		if list[i] != list[lenlist-i-1] {
			return false
		}
	}
	return true
}
