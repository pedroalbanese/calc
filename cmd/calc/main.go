package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	"strconv"
	"github.com/pedroalbanese/calc"
)


	var round = flag.Int("r", 0, "Round to N decimal places.")
	var count = flag.String("e", "", "Mathematical expression.")

func main() {
    flag.Parse()
        if (len(os.Args) < 2) {
	  fmt.Println("Usage of",os.Args[0]+":")
          flag.PrintDefaults()
          os.Exit(1)
        }

	if *round > 0 {
	int64, err := strconv.ParseFloat(calc.Must(calc.Eval(*count)), 64)
		if err != nil {
                        log.Fatal(err)
		}
	var y float64 = float64(int64)
	t := strconv.Itoa(*round)
	var round string = "%."+t+"f"
	fmt.Printf(round, y)
        os.Exit(0)
	} else {
	fmt.Printf(calc.Must(calc.Eval(*count)))
	}
}
