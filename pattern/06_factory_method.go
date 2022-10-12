package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Factory struct {
}

func (f *Factory) Create(fType int) IObjForFactory {
	if fType == 0 {
		return &ObjForFactory1{name: "111", value: 1}
	} else if fType == 1 {
		return &ObjForFactory2{name: "222", value: 2}
	}
	return nil
}

type IObjForFactory interface {
	PrintExample()
}

type ObjForFactory struct {
	name  string
	value int
}

type ObjForFactory1 ObjForFactory

func (o *ObjForFactory1) PrintExample() {
	fmt.Println("Name: ", o.name, " Value: ", o.value)
}

type ObjForFactory2 ObjForFactory

func (o *ObjForFactory2) PrintExample() {
	fmt.Println("Value: ", o.value, " Name: ", o.name)
}
