package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(list []int) *ListNode {
	head := &ListNode{list[0], nil}
	var tmp = head
	for i := 1; i < len(list); i++ {
		tmp.Next = &ListNode{list[i], nil}
		tmp = tmp.Next
	}
	return head
}
