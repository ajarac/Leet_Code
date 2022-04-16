package kth_missing_positive_number

func findKthPositive(arr []int, k int) int {
	p := 0
	i := 0
	counter := 1
	for p < k {
		if i < len(arr) && counter == arr[i] {
			i++
		} else {
			p++
		}
		counter++
	}
	return p + i
}
