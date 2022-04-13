package find_the_distance_value_between_two_arrays

import "testing"

type Input struct {
	arr1 []int
	arr2 []int
	d    int
}

type Test struct {
	input    Input
	expected int
}

func TestFindTheDistanceValue(t *testing.T) {
	var tests = []Test{
		{Input{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2}, 2},
		{Input{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3}, 2},
		{Input{[]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6}, 1},
	}
	for _, test := range tests {
		if output := findTheDistanceValue(test.input.arr1, test.input.arr2, test.input.d); output != test.expected {
			t.Error("TestFindTheDistanceValue Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
