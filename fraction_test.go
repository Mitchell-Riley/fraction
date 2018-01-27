package fraction

import (
	"fmt"
	"testing"
)

func TestBaseRExpansion(t *testing.T) {
	tests := []struct {
		base        int
		numerator   int
		denominator int
		quotient    string
	}{
		{7, 1, 4410, "0.00[003545132261]"},
		{7, 7, 8, "0.[60]"},
		{2, 14, 100, "0.00[10001111010111000010]"},
		{10, 1, 22, "0.0[45]"},
	}

	for _, v := range tests {
		q, _, _, err := BaseRExpansion(v.base, v.numerator, v.denominator)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(q)

		if q != v.quotient {
			fmt.Println("got: ", q)
		}
	}
}
