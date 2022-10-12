package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type ChainOfRespData struct {
	Int int
}

type Chain interface {
	Logic()
	SetNextChain(c *Chain)
}

type ChainExample1 struct {
	ChainOfRespData *ChainOfRespData
	NextChain       Chain
}

func (c *ChainExample1) Logic() {
	fmt.Println("Do ChainExample1 Logic()")
	c.ChainOfRespData.Int++
	if c.NextChain != nil {
		fmt.Println("ChainExample1 Do NextChain")
		c.NextChain.Logic()
	} else {
		fmt.Println("ChainExample1 NextChain = nil")
	}
}
func (c *ChainExample1) SetNextChain(next *Chain) {
	c.NextChain = *next
}

type ChainExample2 struct {
	ChainOfRespData *ChainOfRespData
	NextChain       Chain
}

func (c *ChainExample2) Logic() {
	fmt.Println("Do ChainExample2 Logic()")
	c.ChainOfRespData.Int += 100
	if c.NextChain != nil {
		fmt.Println("ChainExample2 Do NextChain")
		c.NextChain.Logic()
	} else {
		fmt.Println("ChainExample2 NextChain = nil")
	}
}
func (c *ChainExample2) SetNextChain(next *Chain) {
	c.NextChain = *next
}
