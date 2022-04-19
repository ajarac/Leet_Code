package merge_two_sorted_list

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	list1 []int
	list2 []int
}

type Test struct {
	input    Input
	expected []int
}

func TestMergeTwoList(t *testing.T) {
	var tests = []Test{
		{Input{[]int{1, 2, 4}, []int{1, 3, 4}}, []int{1, 1, 2, 3, 4, 4}},
		{Input{[]int{}, []int{}}, []int{}},
		{Input{[]int{}, []int{0}}, []int{0}},
	}
	for _, test := range tests {
		list1 := utils.BuildListNode(test.input.list1)
		list2 := utils.BuildListNode(test.input.list2)
		output := mergeTwoLists(list1, list2)
		if utils.Equal(output.ToArr(), test.expected) {
			if (output != nil && len(test.expected) > 0) || utils.Equal(output.ToArr(), test.expected) == false {
				t.Error("TestMergeTwoList Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
			}
		}
	}
}
