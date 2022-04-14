package sqrt_x_

// 4 -> 2
// 0 1 3 -> 0 1 9
// 2 2 2 -> 4 4 4

// 8 -> 2
// 0 3 5 -> 0 9 25
// 1 2 3 -> 1 4 9

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	left, right := 0, (x)/2+1
	for left+1 < right {
		mid := (left + right) / 2

		quadratic := mid * mid
		if quadratic == x {
			return mid
		}
		qLeft := left * left
		qRight := right * right
		if quadratic < x && x < qRight {
			left = mid
		}
		if qLeft < x && x < quadratic {
			right = mid
		}
	}
	return left
}
