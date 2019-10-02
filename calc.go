package calc

import (
	"math"
	"strconv"
	"strings"
)

// Eval evaluates a mathematical expession in BODMAS order, the
// input string array is assumed to already be tokenized. This method
// panics if the expression is not valid for any reason.
func Eval(s ...string) string {
	odmas := []struct {
		Op    string
		Apply func(lhs, rhs string) string
	}{
		{
			Op:    "^",
			Apply: fop(func(a, b float64) float64 { return math.Pow(a, b) }),
		},
		{
			Op:    "/",
			Apply: fop(func(a, b float64) float64 { return a / b }),
		},
		{
			Op:    "*",
			Apply: fop(func(a, b float64) float64 { return a * b }),
		},
		{
			Op:    "+",
			Apply: fop(func(a, b float64) float64 { return a + b }),
		},
		{
			Op:    "-",
			Apply: fop(func(a, b float64) float64 { return a - b }),
		},
	}

	// evaluate brackets, recursively
	for i := 0; i < len(s); i++ {
		if s[i] == "(" {
			bracketDepth := 0
			for j := i; j < len(s); j++ {
				if s[j] == "(" {
					bracketDepth++
				}
				if s[j] == ")" {
					bracketDepth--
				}
				if s[j] == ")" && bracketDepth == 0 {
					s[i] = Eval(s[i+1 : j]...)
					s = append(s[0:i+1], s[j+1:len(s)]...)
					break
				}
			}
		}
	}

	// apply binary operations in precedence order
	for _, op := range odmas {
		for i := 0; i < len(s); i++ {
			if s[i] == op.Op {
				s[i-1] = op.Apply(s[i-1], s[i+1])
				s = append(s[0:i], s[i+2:len(s)]...)
				i = i - 2
			}
		}
	}

	return strings.Join(s, " ")
}

func fop(fn func(lhs, rhs float64) float64) func(string, string) string {
	return func(a, b string) string {
		lhs, err1 := strconv.ParseFloat(a, 64)
		rhs, err2 := strconv.ParseFloat(b, 64)

		if err1 != nil {
			panic(err1)
		}

		if err2 != nil {
			panic(err2)
		}

		return strconv.FormatFloat(fn(lhs, rhs), 'f', -1, 64)
	}
}
