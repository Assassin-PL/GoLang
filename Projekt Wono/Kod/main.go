package main

import (
	"fmt"
	"strconv"
)

var actorName = "John"

const (
	_ = iota
	a //
)

func main() {
	const MY_CONST int = 1
	const a int16 = 69 // shadowing nastopilo "przykrycie zmiennej a zapisanej w funkcji"

	i := 42
	var n bool = true
	var j int8 = 27
	k := 99
	var m string = "NULL"
	o := 1 == 2 // inicjalizacja zmiennych boolowskich
	m = strconv.Itoa(k)

	printValues(a)
	printValues(MY_CONST)
	printValues(i, j, k)
	printFormattedValues(i, k)
	fmt.Print(actorName + "\n")
	printFormattedValues(m)
	fmt.Print(n, "\n")
	printFormattedValues(o)
	printBitwiseOperations(10, 3)

	// typy zespolone
	var v complex64 = 1 + 2i
	var z complex64 = 4 + 7i
	printFormattedValues(v)
	fmt.Print(v+z, "\n") // dodawanie
	fmt.Print(v-z, "\n") // odejmowanie
	fmt.Print(v*z, "\n") // mnożenie
	fmt.Print(v/z, "\n") // dzielenie

	// typy tekstowe string tak jak w c++
	m = "Dupa"
	fmt.Print(actorName + "  " + m + "  ") // konkatenacja jak w Pythonie, nie ma typu char, bo w stringu są aliasy do liter a nie chary
}

func printValues(values ...interface{}) {
	for _, v := range values {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

func printFormattedValues(values ...interface{}) {
	for _, v := range values {
		fmt.Printf("%v, %T\n", v, v)
	}
}

func printBitwiseOperations(a, b int) {
	fmt.Print(a&b, "\n")  // funkcja boolowska AND
	fmt.Print(a|b, "\n")  // funkcja OR
	fmt.Print(a^b, "\n")  // funkcja XOR
	fmt.Print(a&^b, "\n") // funkcja NAND
	fmt.Print(a<<1, "\n") // przesunięcie bitowe w lewo
	fmt.Print(a>>1, "\n") // przesunięcie bitowe w prawo
}
