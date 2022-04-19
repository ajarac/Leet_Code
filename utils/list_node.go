package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) ToArr() []int {
	if head == nil {
		return []int{}
	}
	var arr []int
	tmp := head
	arr = append(arr, tmp.Val)
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	return arr
}

func BuildListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{list[0], nil}
	var tmp = head
	for i := 1; i < len(list); i++ {
		tmp.Next = &ListNode{list[i], nil}
		tmp = tmp.Next
	}
	return head
}
