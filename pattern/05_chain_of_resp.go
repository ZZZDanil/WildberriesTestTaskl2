package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Data struct {
	int
}

type Chain interface {
	Logic(d *Data)
	NextChain(c *Chain)
}

type ChainExample1 struct {
	Chain *Chain
}

func (c *ChainExample1) Logic(d *Data) {
	// ChainExample1 logic
	d.int++
	(*c.Chain).Logic(d)
}
func (c *ChainExample1) NextChain(next *Chain) {
	c.Chain = next
}

type ChainExample2 struct {
	Chain *Chain
}

func (c *ChainExample2) Logic(d *Data) {
	// ChainExample2 logic
	d.int++
	(*c.Chain).Logic(d)
}
func (c *ChainExample2) NextChain(next *Chain) {
	c.Chain = next
}
