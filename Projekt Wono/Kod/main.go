package main

import (
	"fmt"
)

var JANUSZ string = "TwojaStara"

var (
	a         float32 = 0.09
	actorName string  = "Janusz Kowalski"
)

func main() {
	i := 42
	var j int = 27
	k := 99
	fmt.Println(i, j, k)
	fmt.Printf("%v, %T", i, i)
	fmt.Printf("\n")
	fmt.Printf("%v, %T", a, a)
	fmt.Printf("\n")
	fmt.Print(actorName)
}
