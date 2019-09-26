package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ferhatelmas/pi"
)

func main() {
	n := flag.Int64("n", 1000, "number of digits")
	help := flag.Bool("h", false, "print help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	fmt.Println(pi.Digits(*n))
}
