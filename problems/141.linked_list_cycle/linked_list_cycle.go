package linked_list_cycle

import "leetcode/utils"

func hasCycle(head *utils.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
