package pattern

import (
	"fmt"
	"strconv"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

const (
	t1 = iota
	t2
)

type MyStringBuilder struct {
	value1, value2 string
}

func (b *MyStringBuilder) HardLogic1(buildType int) {
	switch buildType {
	case t1:
		{
			b.value1 = "Build1"
		}
	case t2:
		{
			b.value1 = "Build2"
		}
	default:
		{
			fmt.Println("Error in HardLogic1 buildType : " + strconv.Itoa(buildType))
		}
	}
}

func (b *MyStringBuilder) HardLogic2(buildType int) {
	switch buildType {
	case t1:
		{
			b.value2 = "Build3"
		}
	case t2:
		{
			b.value2 = "Build4"
		}
	default:
		{
			fmt.Println("Error in HardLogic2 buildType : " + strconv.Itoa(buildType))
		}
	}
}
func (b *MyStringBuilder) GetString() string {
	if b.value1 != "" && b.value2 != "" {
		return (b.value1 + " " + b.value2)
	} else {
		return "GetString ERROR: value1: " + b.value1 + " value2: " + b.value2
	}
}
