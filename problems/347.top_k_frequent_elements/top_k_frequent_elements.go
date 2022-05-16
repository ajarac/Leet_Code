package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	mapper := make(map[int]int)
	for _, num := range nums {
		mapper[num]++
	}
	freq := make([][]int, len(nums)+1)

	for n, counter := range mapper {
		freq[counter] = append(freq[counter], n)
	}
	res := make([]int, k)
	cnt := 0
	for i := len(freq) - 1; i > 0 && cnt < k; i-- {
		if freq[i] != nil {
			for _, ii := range freq[i] {
				res[cnt] = ii
				cnt++
			}
		}
	}
	return res
}
