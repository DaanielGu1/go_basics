package main

import (
	"Treino-em-go/tools"
)

func main() {
	list := tools.NewList[int]()
	
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(3)

	// x := tools.NewFindFirstSearcher(tools.NewIntComparable(3))

	// y := tools.NewFindAll(tools.NewIntComparable(3))

}
