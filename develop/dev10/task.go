package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {

	host := "medium.com"
	port := "8080"
	timeOut := time.Second * time.Duration(1)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeOut)
	//conn, err := net.Dial("tcp", "medium.com:12345")
	if err == nil {
		//fmt.Println("----", conn)
		osSignals := make(chan os.Signal, 1)
		listenErr := make(chan error, 1)
		signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

		go req(conn, listenErr, osSignals)
		go resp(conn, listenErr, osSignals)

		select {
		case <-osSignals:
			conn.Close()
		case err = <-listenErr:
			if err != nil {
				log.Fatalf("go-telnet: %s", err)
			}
		}
	} else {

		fmt.Println("err: ", err)

	}

}

func req(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			listenErr <- err
		}

		fmt.Fprintf(conn, text+"\n")
	}
}

func resp(conn net.Conn, listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			listenErr <- err
		}

		fmt.Print(text)
	}
}
