package main

import (
	"fmt"
)

/*
An interface which has zero methods is called empty interface.
 It is represented as interface{}.
 Since the empty interface has zero methods, all types implement the empty interface.
*/

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}
