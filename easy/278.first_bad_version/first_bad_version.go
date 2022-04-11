package first_bad_version

type BadVersion struct {
	version int
}

func firstBadVersion(n int, b *BadVersion) int {
	if n == 1 {
		return 1
	}

	left, right := 1, n

	for right-left > 1 {
		mid := (left + right) / 2
		if b.isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	if b.isBadVersion(left) {
		return left
	} else {
		return right
	}
}

func (b *BadVersion) isBadVersion(n int) bool {
	return n >= b.version
}
