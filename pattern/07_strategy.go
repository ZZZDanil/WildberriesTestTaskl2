package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	DoSome()
}
type Obj struct {
	Strategy
}

func (o *Obj) DoStrategy() {
	o.Strategy.DoSome()
}

type Strategy1 struct {
}

func (s *Strategy1) DoSome() {
	s.Strategy1Func()
}
func (s *Strategy1) Strategy1Func() {
	fmt.Println("Do Strategy1")
}

type Strategy2 struct {
}

func (s *Strategy2) DoSome() {
	s.Strategy2Func()
}
func (s *Strategy2) Strategy2Func() {
	fmt.Println("Do Strategy1")
}
