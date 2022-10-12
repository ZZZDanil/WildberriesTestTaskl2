package main

import (
	"fmt"
	"pattern"
)

func main() {
	GoFacade()
	GoBuilder()
	GoVisitor()
	GoCommand()
	GoChainOfResp()
	GoFactoryMethod()
	GoStrategy()
	GoState()
}

func GoFacade() {
	fmt.Println("<--Facade Pattern-->")

	facade := pattern.CreateFacade()
	facade.GlobalPrint()
	facade.ChangeGlobalPrint("ttt", "ddd")
	facade.GlobalPrint()

	fmt.Println()
}

func GoBuilder() {
	fmt.Println("<--Builder Pattern-->")

	fmt.Println("Start Builder")
	builder := pattern.MyStringBuilder{}
	builder.HardLogic1(1)
	builder.HardLogic2(0)
	result := builder.GetString()
	fmt.Println("Builder Result: ", result)

	fmt.Println("Start New Builder")
	builder = pattern.MyStringBuilder{}
	builder.HardLogic1(100)
	builder.HardLogic2(0)
	result = builder.GetString()
	fmt.Println("New Builder Result: ", result)

	fmt.Println()
}

func GoVisitor() {
	fmt.Println("<--Visitor Pattern-->")

	myVisitor := pattern.MyVisitorRealisation{}
	smallObject := pattern.SmallObjectForVisitor{}
	bigObjectForVisitor := pattern.BigObjectForVisitor{}

	smallObject.Accept(&myVisitor)
	bigObjectForVisitor.Accept(&myVisitor)

	fmt.Println()
}

func GoCommand() {
	fmt.Println("<--Command Pattern-->")

	var invoker pattern.Invoker
	var command pattern.Command
	receiver := pattern.Receiver{}

	command = &pattern.Command1{Receiver: &receiver}
	invoker = &pattern.Invoker1{Command: &command}

	fmt.Println("Receiver int = ", receiver.Int)
	invoker.Invoke()
	fmt.Println("Receiver int = ", receiver.Int)
	fmt.Println()
}

func GoChainOfResp() {
	fmt.Println("<--Chain Of Resp Pattern-->")

	data := pattern.ChainOfRespData{}
	c2 := pattern.ChainExample2{ChainOfRespData: &data}
	c1 := pattern.ChainExample1{ChainOfRespData: &data, NextChain: &c2}

	c1.Logic()

	fmt.Println("Data Value = ", data.Int)
	fmt.Println()
}

func GoFactoryMethod() {
	fmt.Println("<--Factory Method Pattern-->")

	f := pattern.Factory{}
	fmt.Println("Factory(0):")
	f1 := f.Create(0)
	f1.PrintExample()
	fmt.Println("Factory(1):")
	f2 := f.Create(1)
	f2.PrintExample()

	fmt.Println()
}

func GoStrategy() {
	fmt.Println("<--Strategy Pattern-->")

	rootS := pattern.RootStrategy{}
	strategy1 := pattern.StrategyType1{}
	strategy2 := pattern.StrategyType2{}

	fmt.Println("Set Strategy Type 1")
	rootS.Strategy = &strategy1
	rootS.DoStrategy()

	fmt.Println("Set Strategy Type 2")
	rootS.Strategy = &strategy2
	rootS.DoStrategy()

	fmt.Println()
}

func GoState() {
	fmt.Println("<--State Pattern-->")

	stateM := pattern.InitStateObj()
	stateM.Do()
	stateM.Do()
	stateM.Do()
	stateM.Do()
	stateM.Do()

	fmt.Println()
}
