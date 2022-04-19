package merge_two_sorted_list

import "leetcode/utils"

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	var result *utils.ListNode
	var tmp *utils.ListNode
	var next *utils.ListNode
	for list1 != nil || list2 != nil {
		if (list1 == nil && list2 != nil) || (list1 != nil && list2 != nil && list1.Val >= list2.Val) {
			next = list2
			list2 = list2.Next
		} else if (list1 != nil && list2 == nil) || (list1 != nil && list2 != nil && list1.Val < list2.Val) {
			next = list1
			list1 = list1.Next
		}
		if tmp == nil {
			result = next
			tmp = next
		} else {
			tmp.Next = next
			tmp = tmp.Next
		}
	}
	return result
}
