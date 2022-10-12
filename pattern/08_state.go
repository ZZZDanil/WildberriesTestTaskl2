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
type StateMachine struct {
	main, a, b State
}

func InitStateObj() *StateMachine {
	a := &StateA{}
	b := &StateB{}
	return &StateMachine{main: a, a: a, b: b}
}

func (o *StateMachine) Do() {
	o.main.DoSome()
	o.changeState()
}
func (o *StateMachine) changeState() {
	if o.main == o.a {
		o.main = o.b
	} else {
		o.main = o.a
	}
}

type StateA struct {
	StateMachine
}

func (s *StateA) DoSome() {
	fmt.Println("Do StateA")
}

type StateB struct {
	StateMachine
}

func (s *StateB) DoSome() {
	fmt.Println("Do StateB")
	s.main = s.a
}
