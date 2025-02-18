package main

import (
	adder "contain-yourself/swig/adder"
	"fmt"
)

func main() {
	a := adder.NewAdder()
	defer adder.DeleteAdder(a)
	a.Add(1)
	a.Add(2)
	fmt.Printf("value %d\n", a.Get())
}
