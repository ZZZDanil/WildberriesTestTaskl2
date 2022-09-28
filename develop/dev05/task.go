package main

import (
	"bufio"
	"bytes"
	"fmt"
	"main/options"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	// go run task.go -input="example.txt" -pattern="text test" -C 2

	options := options.ReadOptions()
	inFile := openFile(options.Input)
	defer inFile.Close()

	file, findRow := readFile(inFile, options)
	fmt.Println("findRow ", findRow)
	if findRow == -1 {
		fmt.Println("не найдено")
	} else {
		if options.C > 0 {
			PrintRes(file, findRow-options.C, findRow+options.C, options.N)
		} else if options.B > 0 {
			PrintRes(file, findRow-options.B, findRow, options.N)
		} else {
			PrintRes(file, findRow, findRow+options.A, options.N)
		}
	}

}

func openFile(path string) *os.File {
	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return nil
	}

	return inFile
}

func readFile(inFile *os.File, options *options.Options) ([]*[]byte, int) {
	r := bufio.NewReader(inFile)

	file := make([]*[]byte, 0)
	findRow := -1
	basic := func() {
		for i := 0; ; i++ {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}
			file = append(file, &line)
			if findRow < 0 {
				if options.F {
					if options.Pattern == string(line) {
						findRow = i
					}
				} else {
					if res, _ := regexp.Match(options.Pattern, line); res {
						findRow = i
					}
				}
			}
		}
	}
	ignoreCase := func() {
		options.Pattern = strings.ToLower(options.Pattern)
		for i := 0; ; i++ {
			line, _, err := r.ReadLine()
			if err != nil {
				break
			}
			line = bytes.ToLower(line)
			file = append(file, &line)
			if findRow < 0 {
				if options.F {
					if options.Pattern == string(line) {
						findRow = i
					}
				} else {
					if res, _ := regexp.Match(options.Pattern, line); res {
						findRow = i
					}
				}
			}
		}
	}
	if options.I {
		ignoreCase()
	} else {
		basic()
	}
	return file, findRow
}
func PrintRes(data []*[]byte, from, to int, printRow bool) {
	l := len(data)
	to++
	if from < 0 {
		from = 0
	}
	if to > l {
		to = l
	}
	var outFormat func(data ...interface{})
	if !printRow {
		outFormat = func(data ...interface{}) {
			fmt.Println(data[0])
		}
	} else {
		outFormat = func(data ...interface{}) {
			fmt.Println(data[1], " ", data[0])
		}
	}
	for i, d := range data[from:to] {
		outFormat(string(*d), i+from)
	}

}
