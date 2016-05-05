package factorial

import (
	"testing"
)

type TestHelper struct {
	input []int{1,3,7}
	expectations []int{1,6,5040}
}

func TestCalculateFactorial(t *testing.T) {
	helper := new (TestHelper)

	for index, value := range helper.input {
		fn := CalculateFactorial(value)

		if val := fn(); val != helper.ex {
			t.Fatalf("Expected %d, got %", helper)
		}
	}

}
