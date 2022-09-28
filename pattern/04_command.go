package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute()
}
type Receiver struct {
	int
}
type Invoker interface {
	Invoke()
}

/////////////////////////

type Command1 struct {
	Receiver
}

func (c *Command1) Execute() {
	c.Receiver.int++
}

type Invoker1 struct {
	Command
}

func (i *Invoker1) Invoke() {
	i.Command.Execute()
}
