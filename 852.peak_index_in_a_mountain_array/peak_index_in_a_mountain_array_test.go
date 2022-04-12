package peak_index_in_a_mountain_array

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
	}

	for _, test := range tests {
		if output := peakIndexInMountainArray(test.input); output != test.expected {
			t.Error("TestPeakIndexInMountainArray Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
