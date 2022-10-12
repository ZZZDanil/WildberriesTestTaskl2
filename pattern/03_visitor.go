package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type MyVisitor interface {
	VisitSmallObject(s *SmallObjectForVisitor)
	VisitBigObject(b *BigObjectForVisitor)
}

// структура визитёра, содержащая логику общего типа для нескольких других структур
type MyVisitorRealisation struct{}

func (m *MyVisitorRealisation) VisitSmallObject(s *SmallObjectForVisitor) {
	fmt.Println("MyVisitorRealisation: Visit SmallObject")
}
func (m *MyVisitorRealisation) VisitBigObject(b *BigObjectForVisitor) {
	fmt.Println("MyVisitorRealisation: Visit BigObject")

}

// простая структура
type SmallObjectForVisitor struct {
}

func (s *SmallObjectForVisitor) Accept(m MyVisitor) {
	m.VisitSmallObject(s)
}

// очень *сложная структура
type BigObjectForVisitor struct {
}

func (b *BigObjectForVisitor) Accept(m MyVisitor) {
	m.VisitBigObject(b)
}
