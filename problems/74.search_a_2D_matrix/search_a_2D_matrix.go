package search_a_2D_matrix

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	start := 0
	end := m*n - 1
	for start <= end {
		mid := (start + end) / 2
		row := mid / n
		col := mid % n
		midVal := matrix[row][col]
		if midVal == target {
			return true
		} else if midVal > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}
