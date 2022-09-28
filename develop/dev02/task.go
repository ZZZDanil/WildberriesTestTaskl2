package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var inputString string

	// flags declaration using flag package
	flag.StringVar(&inputString, "u", "ppppppp", "String for parsing")
	flag.Parse()

	fmt.Println("input:  ", inputString)
	s, _ := stringPars(inputString)
	fmt.Println("output: ", s)
}

func stringPars(inputString string) (string, error) {
	if inputString == "" {
		return "", nil
	} else {

		out := make([]rune, 0, 10)
		var runeBuf rune
		isEscape := false
		isEscapeSimbol := false
		isCountRep := false
		for _, r := range inputString {
			if r == '\\' {
				if isEscape == false {
					isEscape = true
					isEscapeSimbol = true
					continue
				} else {
					return "некорректная строка", errors.New("\\\\s")
				}
			} else {
				isEscape = false
				value, err := strconv.Atoi(string(r))
				if err != nil || isEscapeSimbol {
					isEscapeSimbol = false
					isCountRep = false
					runeBuf = r
					out = append(out, r)
					continue
				} else {
					if !isCountRep {
						isCountRep = true
						o := make([]rune, value-1)
						for i := range o {
							o[i] = runeBuf
						}
						out = append(out, o...)
					} else {
						return "некорректная строка", errors.New("два числа подряд")
					}
				}

			}

		}
		return string(out), nil
	}
}
