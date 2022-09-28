package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor interface {
	SomeLogicForBasicStruct(b *BasicStruct)
}

type VisitorRealisation1 struct {
}

func (v *VisitorRealisation1) SomeLogicForBasicStruct(b *BasicStruct) {
	fmt.Println(b.someData)
}

type VisitorRealisation2 struct {
}

func (v *VisitorRealisation2) SomeLogicForBasicStruct(b *BasicStruct) {
	fmt.Println("VisitorRealisation2 ", b.someData)
}

type BasicStruct struct {
	someData string
}
