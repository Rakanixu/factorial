package factorial

import (
	"testing"
)

type TestHelper struct {
	input []int
	expectations []int
}

func TestCalculateFactorial(t *testing.T) {
	helper := TestHelper {
		input: []int {1, 3, 7},
		expectations: []int {1, 6, 5040},
	}

	for index, value := range helper.input {
		result := CalculateFactorial(value)

		if result != helper.expectations[index] {
			t.Fatalf("Expected %d, got %d", helper.expectations[index], result)
		}
	}

}
