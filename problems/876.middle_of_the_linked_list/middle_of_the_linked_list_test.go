package middle_of_the_linked_list

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	head     *utils.ListNode
	expected *utils.ListNode
}

func TestMiddleNode(t *testing.T) {
	var tests = []Test{
		{utils.BuildListNode([]int{1, 2, 3, 4, 5}), utils.BuildListNode([]int{3, 4, 5})},
		{utils.BuildListNode([]int{1, 2, 3, 4, 5, 6}), utils.BuildListNode([]int{4, 5, 6})},
	}

	for _, test := range tests {
		if output := middleNode(test.head); output.Val != test.expected.Val {
			t.Error("TestMiddleNode Failed: {} inputted, {} expected, received: {}", test.head, test.expected, output)
		}
	}
}
