package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	o := ReadOptions()
	columns := make([]int, 0)
	for _, num := range strings.Split(o.F, ",") {
		i, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		columns = append(columns, i)
	}
	inputData := make([]string, 0)
	breakInput := make(chan os.Signal)
	signal.Notify(breakInput, os.Interrupt)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input:")
	for scanner.Scan() {

		select {
		case <-breakInput:
			{
				inputData = append(inputData, scanner.Text())
				break
			}
		default:
			{

				inputData = append(inputData, scanner.Text())
			}
		}

	}
	fmt.Print("Cut:")
	for _, s := range inputData {
		words := strings.Split(s, o.D)
		printWord := func(words []string, indexes []int) {
			for _, ind := range indexes {
				fmt.Print(words[ind])
			}
			fmt.Print("\n")
		}
		if o.S && len(words) > 1 || !o.S {
			printWord(words, columns)
		}
	}
}

type Options struct {
	F, D string
	S    bool
}

func ReadOptions() (options *Options) {
	options = new(Options)
	flag.StringVar(&options.F, "f", "", "fields")
	flag.StringVar(&options.D, "d", "	", "delimiter")
	flag.BoolVar(&options.S, "s", false, "separated")
	flag.Parse()

	if options.F == "" {
		fmt.Println("не указаны колонки")
		os.Exit(-1)
	}

	return
}
