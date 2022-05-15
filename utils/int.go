package utils

func IntMaxValue() int {
	maxUint := ^uint32(0)
	return int(maxUint >> 1)
}
