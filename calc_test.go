package calc_test

import (
	"testing"

	"github.com/julz/calc"
	"gotest.tools/assert"
)

func TestEval(t *testing.T) {
	examples := []struct {
		Title  string
		Input  string
		Expect string
	}{
		{
			Title:  "Single integer",
			Input:  "2",
			Expect: "2",
		},
		{
			Title:  "Simple Addition",
			Input:  "2+3",
			Expect: "5",
		},
		{
			Title:  "Simple Addition (with Whitespace)",
			Input:  "2 + 3",
			Expect: "5",
		},
		{
			Title:  "Simple Exponent",
			Input:  "2^3",
			Expect: "8",
		},
		{
			Title:  "Simple Division",
			Input:  "8/2",
			Expect: "4",
		},
		{
			Title:  "Modulus",
			Input:  "8%3",
			Expect: "2",
		},
		{
			Title:  "Fractional Division",
			Input:  "2/8",
			Expect: "0.25",
		},
		{
			Title:  "Complex Addition",
			Input:  "2+3+5+9",
			Expect: "19",
		},
		{
			Title:  "Addition and Multiplication (in textual order)",
			Input:  "2*3+5+9",
			Expect: "20",
		},
		{
			Title:  "Addition and Multiplication (multiplication applied before division)",
			Input:  "5+9+2*3",
			Expect: "20",
		},
		{
			Title:  "arbitrary whitespace",
			Input:  "  5 +9  +2* 3",
			Expect: "20",
		},
		{
			Title:  "Single Brackets",
			Input:  "6+(5+2)*3",
			Expect: "27",
		},
		{
			Title:  "Multiple Brackets",
			Input:  "6+(5+2)*(3-1)",
			Expect: "20",
		},
		{
			Title:  "Overlapping Brackets",
			Input:  "6+(5+(8/2)+7)*(3-1)",
			Expect: "38",
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			result, err := calc.Eval(eg.Input)
			assert.NilError(t, err)
			assert.Equal(t, eg.Expect, result)
		})
	}
}

func TestEvalErrors(t *testing.T) {
	examples := []struct {
		Title             string
		Input             string
		ExpectErrContains string
	}{
		{
			Title:             "Unknown char",
			Input:             "# + 3",
			ExpectErrContains: "expected number got '#'",
		},
		{
			Title:             "Unbalanced brackets",
			Input:             "1 + 2 + (3",
			ExpectErrContains: "mismatched brackets",
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			_, err := calc.Eval(eg.Input)
			assert.ErrorContains(t, err, eg.ExpectErrContains)
		})
	}
}
