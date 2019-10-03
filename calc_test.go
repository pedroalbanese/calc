package calc_test

import (
	"testing"

	"github.com/julz/calc"
	"gotest.tools/assert"
)

func TestParse(t *testing.T) {
	examples := []struct {
		Title  string
		Input  []string
		Expect string
	}{
		{
			Title:  "Single integer",
			Input:  []string{"2"},
			Expect: "2",
		},
		{
			Title:  "Simple Addition",
			Input:  []string{"2", "+", "3"},
			Expect: "5",
		},
		{
			Title:  "Simple Exponent",
			Input:  []string{"2", "^", "3"},
			Expect: "8",
		},
		{
			Title:  "Simple Division",
			Input:  []string{"8", "/", "2"},
			Expect: "4",
		},
		{
			Title:  "Fractional Division",
			Input:  []string{"2", "/", "8"},
			Expect: "0.25",
		},
		{
			Title:  "Complex Addition",
			Input:  []string{"2", "+", "3", "+", "5", "+", "9"},
			Expect: "19",
		},
		{
			Title:  "Addition and Multiplication (in textual order)",
			Input:  []string{"2", "*", "3", "+", "5", "+", "9"},
			Expect: "20",
		},
		{
			Title:  "Addition and Multiplication (multiplication applied before division)",
			Input:  []string{"5", "+", "9", "+", "2", "*", "3"},
			Expect: "20",
		},
		{
			Title:  "Single Brackets",
			Input:  []string{"6", "+", "(", "5", "+", "2", ")", "*", "3"},
			Expect: "27",
		},
		{
			Title:  "Multiple Brackets",
			Input:  []string{"6", "+", "(", "5", "+", "2", ")", "*", "(", "3", "-", "1", ")"},
			Expect: "20",
		},
		{
			Title:  "Overlapping Brackets",
			Input:  []string{"6", "+", "(", "5", "+", "(", "8", "/", "2", ")", "+", "7", ")", "*", "(", "3", "-", "1", ")"},
			Expect: "38",
		},
	}

	for _, eg := range examples {
		t.Run(eg.Title, func(t *testing.T) {
			result := calc.EvalTokens(eg.Input...)
			assert.Equal(t, eg.Expect, result)
		})
	}
}
