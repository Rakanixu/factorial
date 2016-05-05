package factorial

func CalculateFactorial(factorial int) int {
	if (factorial > 0) {
		return CalculateFactorial(factorial - 1) * factorial
	}

	return 1
}