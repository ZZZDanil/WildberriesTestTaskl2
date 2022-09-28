package options

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Input, Pattern string
	A, B, C, Count int
	I, V, F, N     bool
}

func ReadOptions() (options *Options) {
	options = new(Options)
	flag.StringVar(&options.Input, "input", "", "Input")
	flag.StringVar(&options.Pattern, "pattern", "", "Pattern")
	flag.IntVar(&options.A, "A", 3, "after")
	flag.IntVar(&options.B, "B", 0, "before")
	flag.IntVar(&options.C, "C", 0, "context")
	flag.IntVar(&options.Count, "c", 0, "count")
	flag.BoolVar(&options.I, "i", false, "ignore-case")
	flag.BoolVar(&options.V, "v", false, "invert")
	flag.BoolVar(&options.F, "F", false, "fixed")
	flag.BoolVar(&options.N, "n", false, "line num")
	flag.Parse()

	if options.Input == "" {
		fmt.Println("не указан файл")
		os.Exit(-1)
	}
	if options.C > 0 {
		options.A = 0
		options.B = 0
	}
	if options.A < 0 {
		options.A = 0
	}
	if options.B < 0 {
		options.B = 0
	}
	if options.C < 0 {
		options.C = 0
	}
	return
}
