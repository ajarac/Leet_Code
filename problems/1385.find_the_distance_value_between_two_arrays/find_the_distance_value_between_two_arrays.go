package find_the_distance_value_between_two_arrays

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	counter := 0
	for _, num1 := range arr1 {
		isOk := true
		for _, num2 := range arr2 {
			if abs(num1-num2) <= d {
				isOk = false
				break
			}
		}
		if isOk {
			counter++
		}
	}

	return counter
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
