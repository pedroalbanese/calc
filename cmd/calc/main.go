package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"

	"github.com/julz/calc"
)

func main() {
	var s scanner.Scanner
	s.Init(strings.NewReader(os.Args[1]))

	var r []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		r = append(r, s.TokenText())
	}

	fmt.Println(calc.EvalTokens(r...))
}
