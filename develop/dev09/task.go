package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// go run task.go -site https://ya.ru/ -out test.html
	WGet(ReadOptions())
}

type Options struct {
	Site, OutFilePath string
}

func ReadOptions() (options *Options) {
	options = new(Options)
	flag.StringVar(&options.Site, "site", "", "Site")
	flag.StringVar(&options.OutFilePath, "out", "OutFile.html", "OutFilePath")
	flag.Parse()

	if options.Site == "" {
		fmt.Println("не указан сайт")
		os.Exit(-1)
	}

	return
}

func WGet(o *Options) {
	resp, err := http.Get((*o).Site)
	if err == nil {
		file, err := os.Create((*o).OutFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(-2)
		}
		io.Copy(file, resp.Body)
	} else {
		fmt.Println(err)
		os.Exit(-3)
	}
}
