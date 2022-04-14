package utils

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func DeepEqual(mat1 [][]int, mat2 [][]int) bool {
	if len(mat1) != len(mat2) {
		return false
	}
	for i := range mat1 {
		if Equal(mat1[i], mat2[i]) == false {
			return false
		}
	}
	return true
}
