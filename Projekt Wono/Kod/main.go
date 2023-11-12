package main

import (
	"fmt"
	"strconv"
)

var JANUSZ string = "TwojaStara"

var (
	a         float32 = 0.09
	actorName string  = "Janusz Kowalski"
)

func main() {
	i := 42
	var n bool = true
	var j int = 27
	k := 99
	var m string = "NULL"
	o := 1 == 2 // inicjalizacja zmiennych boolowskich
	m = strconv.Itoa(k)
	fmt.Println(i, j, k)
	fmt.Printf("%v, %T", i, i)
	fmt.Printf("\n")
	fmt.Printf("%v, %T", a, a)
	fmt.Printf("\n")
	fmt.Printf("%v, %T", k, k)
	fmt.Printf("\n")
	fmt.Print(actorName)
	fmt.Printf("\n")
	fmt.Printf("%v, %T", m, m)
	fmt.Print(n)
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("%v, %T", o, o)
}
