package valid_perfect_square

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 0, (num)/2+1
	for left+1 < right {
		mid := (left + right) / 2
		quadratic := mid * mid
		if quadratic == num {
			return true
		}
		if quadratic < num {
			left = mid
		} else {
			right = mid
		}
	}
	return false
}
