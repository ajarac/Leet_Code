package guess_number_higher_or_lower

type GuessGame struct {
	pick int
}

func guessNumberHigherOrLower(n int, g *GuessGame) int {
	if n == 1 {
		return 1
	}

	left, right := 1, n

	for right-left > 1 {
		mid := (left + right) / 2
		switch g.guess(mid) {
		case -1:
			right = mid
			break
		case 1:
			left = mid
			break
		default:
			return mid
		}
	}

	if g.guess(left) == 0 {
		return left
	} else {
		return right
	}
}

func (g *GuessGame) guess(n int) int {
	switch {
	case n > g.pick:
		return -1
	case n < g.pick:
		return 1
	default:
		return 0
	}
}
