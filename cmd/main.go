package main

import (
	"fmt"
	"github.com/hbyscx001/Fuck-Alg/data-structure"
)

func main() {
	a := datastrucure.NewMinIntHeap()
	a.Insert(1)
	a.Insert(0)
	a.Insert(100)
	a.Insert(10)

	if v, ok := a.Pop(); ok {
		fmt.Println(v)
	}

	if v, ok := a.Pop(); ok {
		fmt.Println(v)
	}

	fmt.Println(a.Data)
}