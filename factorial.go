package factorial

func CalculateFactorial(factorial int64) int64 {
	if factorial > 0 {
		return CalculateFactorial(factorial-1) * factorial
	}

	return 1
}
