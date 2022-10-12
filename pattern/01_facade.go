package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Facade struct {
	BackLogic1
	BackLogic2
}

func CreateFacade() Facade {
	f := Facade{}

	f.BackLogic1.Back1()
	f.BackLogic2.Back2()
	return f
}

// доступна только простая реализация
func (f *Facade) GlobalPrint() {

	fmt.Println("Простая функция Фасада GlobalPrint:\n\tBackLogic1: ", f.BackLogic1.GetData(), " BackLogic2: ", f.BackLogic2.GetData())
}

func (f *Facade) ChangeGlobalPrint(data1, data2 string) {
	fmt.Println("Простая функция Фасада ChangeGlobalPrint")
	f.BackLogic1.SetData(data1)
	f.BackLogic2.SetData(data2)
}

// ******************************************************
// иная сложная реализация
type BackLogic1 struct {
	data string
}

func (b *BackLogic1) Back1() {
	fmt.Println("Сложные Действия BackLogic1 (data = Back1)")
	b.data = "Back1"
}
func (b *BackLogic1) GetData() string {
	fmt.Println("Сложные Действия BackLogic1 (GetData)")
	return b.data
}
func (b *BackLogic1) SetData(s string) {
	fmt.Println("Сложные Действия BackLogic1 (SetData)")
	b.data = s
}

type BackLogic2 struct {
	data string
}

func (b *BackLogic2) Back2() {
	fmt.Println("Сложные Действия BackLogic2 (data = Back2)")
	b.data = "Back2"
}
func (b *BackLogic2) GetData() string {
	fmt.Println("Сложные Действия BackLogic2 (GetData)")
	return b.data
}
func (b *BackLogic2) SetData(s string) {
	fmt.Println("Сложные Действия BackLogic2 (SetData)")
	b.data = s
}
