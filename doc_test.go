package calc_test

import (
	"fmt"

	"github.com/julz/calc"
)

func Example() {
	fmt.Println(calc.Must(calc.Eval("2 + (3 * 4) / 2")))
	// Output: 8
}
