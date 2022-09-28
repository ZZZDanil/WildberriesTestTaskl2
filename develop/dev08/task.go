package main

import (
	"os"
	"os/exec"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// go build
	// run .exe
	e := exec.Command("powershell")

	e.Stderr = os.Stderr
	e.Stdout = os.Stdout
	e.Stdin = os.Stdin
	e.Run()

}
