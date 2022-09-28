package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
type State interface {
	DoSome()
}
type Obj struct {
	main, a, b State
}

func (o *Obj) Do() {
	o.main.DoSome()
}
func (o *Obj) ChangeStateToA() {
	o.main = o.a
}
func (o *Obj) ChangeStateToB() {
	o.main = o.b
}

type StateA struct {
	Obj
}

func (s *StateA) DoSome() {
	fmt.Println("Do StateA")
}

type StateB struct {
	Obj
}

func (s *StateB) DoSome() {
	fmt.Println("Do StateB")
	s.main = s.a
}
