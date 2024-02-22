package factorial

func calculate(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculate(n-1)
}
