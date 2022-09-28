package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Factory struct {
}

func (f *Factory) Create(fType int) *Obj {
	if fType == 0 {
		return &Obj{name: "111", value: 1}
	} else if fType == 1 {
		return &Obj{name: "222", value: 2}
	}
	return nil
}

type Obj struct {
	name  string
	value int
}
