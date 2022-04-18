package middle_of_the_linked_list

import "leetcode/utils"

func middleNode(head *utils.ListNode) *utils.ListNode {
	sp, fp := head, head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}

	return sp
}
