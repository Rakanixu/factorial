package factorial

func CalculateFactorial(factorial int32) int32 {
	if factorial > 0 {
		return CalculateFactorial(factorial-1) * factorial
	}

	return 1
}
