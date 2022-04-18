package linked_list_cycle

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	input    *utils.ListNode
	pos      int
	expected bool
}

func TestHasCycle(t *testing.T) {
	var tests = []Test{
		{utils.BuildListNode([]int{3, 2, 0, -4}), 1, true},
		{utils.BuildListNode([]int{1, 2}), 0, true},
		{utils.BuildListNode([]int{1}), -1, false},
	}

	for _, test := range tests {
		addCycle(test.input, test.pos)
		if output := hasCycle(test.input); output != test.expected {
			t.Error("TestHasCycle Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func addCycle(head *utils.ListNode, pos int) {
	tmp := head
	counter := 0
	var cycle *utils.ListNode = nil
	var last *utils.ListNode
	for tmp != nil {
		if counter == pos {
			cycle = tmp
		}
		last = tmp
		tmp = tmp.Next
		counter++
	}
	last.Next = cycle
}
