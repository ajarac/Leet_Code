package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	sp, fp := head, head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}

	return sp
}
