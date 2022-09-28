package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var options *Options

func init() {
	options = GetOptions()
}

// go run task.go -in=C:/Users/ZZDan/Desktop/uuu.txt -out=C:/Users/ZZDan/Desktop/out_uuu.txt -k=0 -u=true -r=true
func main() {
	//options = GetOptions()
	// путь указан
	if options.in != "" {
		if _, err := os.Stat(options.in); err == nil {
			data, _ := os.ReadFile(options.in)
			if len(data) > 0 {
				data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
				rows := bytes.Split(data, []byte("\n"))
				sorted := mySort(&rows)
				os.WriteFile(options.out, sorted, 0644)
			} else {
				fmt.Println("Файл пуст")
				os.Exit(0)
			}
		} else {
			fmt.Println("Файл не найден")
			os.Exit(-1)
		}
	}
}

func mySort(rows *[][]byte) []byte {
	fmt.Println(options)
	sortColimns := make([]SortColimn, len(*rows))
	for i, v := range *rows {
		words := bytes.Split(v, []byte(" "))
		sortColimns[i] = *CreateSortColimn(&words[options.column], v, options.isValueSort)
	}
	if options.isReversSort {
		sort.Sort(SortColimnsRevers(sortColimns))
	} else {
		sort.Sort(SortColimns(sortColimns))
	}
	if options.isUniqueSort {
		var buf []byte
		for i := 0; i < len(sortColimns); i++ {
			if reflect.DeepEqual(sortColimns[i].row, buf) && i < len(sortColimns) {
				sortColimns = append(sortColimns[0:i], sortColimns[i+1:]...)
				i--
			} else {
				buf = sortColimns[i].row
			}
		}
	}
	out := make([]byte, 0)
	for _, v := range sortColimns {
		out = append(out[:], v.row...)
		out = append(out[:], []byte("\n")...)
	}
	return out
}

func GetOptions() *Options {
	options := new(Options)
	flag.StringVar(&options.in, "in", "", "Column")
	flag.StringVar(&options.out, "out", "out.txt", "Column")
	flag.IntVar(&options.column, "k", 0, "Column")
	flag.BoolVar(&options.isValueSort, "n", false, "Column")
	flag.BoolVar(&options.isReversSort, "r", false, "Column")
	flag.BoolVar(&options.isUniqueSort, "u", false, "Column")
	flag.Parse()
	return options
}

type Options struct {
	in           string
	out          string
	column       int
	isValueSort  bool
	isReversSort bool
	isUniqueSort bool
}
type SortColimn struct {
	value interface{}
	row   []byte
}

func CreateSortColimn(word *[]byte, row []byte, isValueSort bool) *SortColimn {
	s := string(*word)

	if isValueSort {
		i, _ := strconv.Atoi(s)
		return &SortColimn{value: i, row: row}
	} else {
		return &SortColimn{value: s, row: row}
	}

}

type SortColimns []SortColimn
type SortColimnsRevers []SortColimn

func Less(i, j SortColimn, isRevers bool) bool {
	switch i.value.(type) {
	case int:
		{
			switch j.value.(type) {
			case int:
				{
					if isRevers {
						return i.value.(int) > j.value.(int)
					} else {
						return i.value.(int) < j.value.(int)
					}
				}
			default:
				if isRevers {
					return strconv.Itoa(i.value.(int)) > j.value.(string)
				} else {
					return strconv.Itoa(i.value.(int)) < j.value.(string)
				}
			}
		}
	default:
		switch j.value.(type) {
		case int:
			{
				if isRevers {
					return i.value.(string) > strconv.Itoa(j.value.(int))
				} else {
					return i.value.(string) < strconv.Itoa(j.value.(int))
				}
			}
		default:
			if isRevers {
				return i.value.(string) > j.value.(string)
			} else {
				return i.value.(string) < j.value.(string)
			}
		}
	}

}

func (s SortColimns) Len() int {
	return len(s)
}

func (s SortColimns) Less(i, j int) bool {
	return Less(s[i], s[j], false)
}

func (s SortColimns) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortColimnsRevers) Len() int {
	return len(s)
}

func (s SortColimnsRevers) Less(i, j int) bool {
	return Less(s[i], s[j], true)
}

func (s SortColimnsRevers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
