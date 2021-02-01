package calculator_test

import (
	"calculator"
	"testing"
)

type squareTestCase struct {
	n, want float64
}

type divideTestCase struct {
	x, y, want float64
	errExpected bool `default: false`
}

func TestSquare(t *testing.T) {
	t.Parallel()

	var testCases = []squareTestCase {
		{ n: 0, want: 0 },
		{ n: 2.0, want: 4.0 },
		{ n: -2.0, want: 4.0 },
	}

	for _, tc := range testCases {
		got := calculator.Square(tc.n)
		if tc.want != got {
			t.Errorf("Square(%f): want %f, got %f", tc.n, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	var testCases = []divideTestCase {
		{ x: 2.0, y: 1.0, want: 2.0, errExpected: false, },
		{ x: 2.0, y: 0, want: 0, errExpected: true, },
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.x, tc.y)
		if !tc.errExpected != (err != nil) && tc.want != got {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.x, tc.y, err != nil)
		}
	}
}
