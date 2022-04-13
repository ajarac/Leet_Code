package reverse_integer

func reverse(x int) int {
	rev := 0
	for x != 0 {
		digit := x % 10
		rev = rev*10 + digit
		x = x / 10
	}
	maxUint := ^uint32(0)
	maxInt := int(maxUint >> 1)
	minInt := -maxInt
	if rev > maxInt || rev < minInt {
		return 0
	}
	return rev
}
