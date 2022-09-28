package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// доступна только простая реализация
type Facade struct {
	BackLogic1
	BackLogic2
}

func (f *Facade) CreateFacade(data1, data2 string) {
	f.BackLogic1 = BackLogic1{}
	f.BackLogic2 = BackLogic2{}

	f.BackLogic1.Back1(&data1)
	f.BackLogic2.Back2(&data2)
}

// всё сложное - скрыто
type BackLogic1 struct {
	data string
}

func (b *BackLogic1) Back1(data *string) {
	fmt.Println(*data)
}

type BackLogic2 struct {
	data string
}

func (b *BackLogic2) Back2(data *string) {
	fmt.Println(*data)
}
