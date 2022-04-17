package middle_of_the_linked_list

import "testing"

type Test struct {
	head     *ListNode
	expected *ListNode
}

func TestMiddleNode(t *testing.T) {
	var tests = []Test{
		{buildListNode([]int{1, 2, 3, 4, 5}), buildListNode([]int{3, 4, 5})},
		{buildListNode([]int{1, 2, 3, 4, 5, 6}), buildListNode([]int{4, 5, 6})},
	}

	for _, test := range tests {
		if output := middleNode(test.head); output.Val != test.expected.Val {
			t.Error("TestMiddleNode Failed: {} inputted, {} expected, received: {}", test.head, test.expected, output)
		}
	}
}

func buildListNode(list []int) *ListNode {
	head := &ListNode{list[0], nil}
	var tmp = head
	for i := 1; i < len(list); i++ {
		tmp.Next = &ListNode{list[i], nil}
		tmp = tmp.Next
	}
	return head
}
