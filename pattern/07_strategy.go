package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	doSome()
}
type RootStrategy struct {
	Strategy
}

func (o *RootStrategy) DoStrategy() {
	o.Strategy.doSome()
}

type StrategyType1 struct {
}

func (s *StrategyType1) doSome() {
	s.strategy1Func()
}
func (s *StrategyType1) strategy1Func() {
	fmt.Println("Do Strategy1")
}

type StrategyType2 struct {
}

func (s *StrategyType2) doSome() {
	s.strategy2Func()
}
func (s *StrategyType2) strategy2Func() {
	fmt.Println("Do Strategy2")
}
