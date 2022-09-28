package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// https://socketloop.com/tutorials/golang-get-current-time-from-the-internet-time-server-ntp-example
	ntpTime, err := ntp.Time("time.apple.com,")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Local time: ", time.Now().Local().Format(time.UnixDate))
	fmt.Println("NTP time:   ", ntpTime.Format(time.UnixDate))
	os.Exit(1)

}
