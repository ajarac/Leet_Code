package arranging_coins

import "math"

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	root1, root2 := QuadraticEquation(1.0, 1.0, float64(-2.0*n))
	return int(math.Max(root1, root2))
}

func QuadraticEquation(a float64, b float64, c float64) (float64, float64) {
	discriminant := math.Sqrt(b*b - (4.0 * a * c))
	firstRoot := (-b - discriminant) / (2 * a)
	secondRoot := (-b + discriminant) / (2 * a)
	return firstRoot, secondRoot
}
