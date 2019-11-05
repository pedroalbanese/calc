package calc_test

import (
	"fmt"

	"github.com/julz/calc"
)

func ExampleEval() {
	fmt.Println(calc.Must(calc.Eval("2 + (3 * 4) / 2")))
	// Output: 8
}

func ExampleEvalVars() {
	fmt.Println(calc.Must(calc.EvalVars("(2 + (var1 * 4) / var2) * floatvar", map[string]interface{}{
		"var1":     3,
		"var2":     2,
		"floatvar": 0.5,
	})))
	// Output: 4
}
