package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

const (
	t1 = iota
	t2
)

type Builder struct {
}

func (b *Builder) Build(buildType int) {
	switch buildType {
	case t1:
		{
			Build1()
		}
	case t2:
		{
			Build2()
		}
	}
}

func Build1() {
	fmt.Println("Build1")
}
func Build2() {
	fmt.Println("Build2")
}
