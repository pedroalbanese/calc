// Package calc contains methods for evaluating mathematical expressions in string format
package calc

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/scanner"
	"log"
)

// Must returns the `result` parameter unless `err` is non-nil.
// If `err` is non-nil, `Must` panics.
func Must(result string, err error) string {
	if err != nil {
//		panic(err)
                log.Fatal(err)
	}

	return result
}

// Eval evaluates a mathematical expession in BODMAS order.
func Eval(expr string) (string, error) {
	return EvalVars(expr, make(map[string]interface{}))
}

// EvalVars evaluates a mathematical expression in BODMAS order. Any variables
// in the expression are replaced with the corresponding entry in the `vars`
// parameter before evaluation.
//
// Currently, all values in the `vars` array are expected to be a type
// that can be passed to `fmt.Sprintf("%d")`
func EvalVars(expr string, vars map[string]interface{}) (string, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(expr))

	var r []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		token := s.TokenText()
		if v, ok := vars[token]; ok {
			switch v.(type) {
			case int:
				token = fmt.Sprintf("%d", v)
			case int64:
				token = fmt.Sprintf("%d", v)
			case float64:
				token = fmt.Sprintf("%f", v)
			default:
				return "", fmt.Errorf("unsupported var type: %T", v)
			}
		}

		r = append(r, token)
	}

	return evalTokens(r...)
}

// evalTokens evaluates a mathematical expression in BODMAS order.
// The input is assumed to already be properly tokenized.
func evalTokens(s ...string) (string, error) {
	var bodmas = []op{
		evalBrackets,
		binaryOp("^", func(a, b float64) float64 { return math.Pow(a, b) }),
		binaryOp("/", func(a, b float64) float64 { return a / b }),
		binaryOp("*", func(a, b float64) float64 { return a * b }),
		binaryOp("+", func(a, b float64) float64 { return a + b }),
		binaryOp("-", func(a, b float64) float64 { return a - b }),
		binaryOp("%", func(a, b float64) float64 { return float64(int(a) % int(b)) }),
	}

	var err error
	for _, op := range bodmas {
		s, err = op(s)
		if err != nil {
			return "", err
		}
	}

	return strings.Join(s, " "), nil
}

// evalBrackets recursively evaluates the bracketed expressions
// in a stream of tokens. The return value is the top-level expression, with
// any bracketed sub-expressions replaced with their evaluated result.
func evalBrackets(s []string) ([]string, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == "(" {
			bracketDepth := 0
			for j := i; j < len(s); j++ {
				switch s[j] {
				case "(":
					bracketDepth++
				case ")":
					bracketDepth--
				}

				if s[j] == ")" && bracketDepth == 0 {
					bracketResult, err := evalTokens(s[i+1 : j]...)
					if err != nil {
						return nil, err
					}

					s[i] = bracketResult
					s = append(s[0:i+1], s[j+1:]...)
					break
				}

				// if we get to the end of the tokens and we never found
				// a matching bracket, the expression is invalid
				if j == len(s)-1 {
					return nil, errors.New("mismatched brackets, expected to find ')' but reached end of tokens")
				}
			}
		}
	}

	return s, nil
}

type op func([]string) ([]string, error)

// binaryOp returns a function that applies a binary operation to a stream of tokens
func binaryOp(symbol string, fn func(float64, float64) float64) op {
	return func(s []string) ([]string, error) {
		for i := 0; i < len(s); i++ {
			if s[i] == symbol {
				lhs, err := strconv.ParseFloat(s[i-1], 64)
				if err != nil {
					return nil, fmt.Errorf("expected number got '%s': %s", s[i-1], err)
				}

				rhs, err := strconv.ParseFloat(s[i+1], 64)
				if err != nil {
					return nil, fmt.Errorf("expected number got '%s': %s", s[i+1], err)
				}

				s[i-1] = strconv.FormatFloat(fn(lhs, rhs), 'f', -1, 64)

				s = append(s[0:i], s[i+2:]...)
				i = i - 2
			}
		}

		return s, nil
	}
}
