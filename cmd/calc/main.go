package main

import (
	"fmt"
	"os"

	"github.com/pedroalbanese/calc"
)

func main() {
	fmt.Println(calc.Must(calc.Eval(os.Args[1])))
}
