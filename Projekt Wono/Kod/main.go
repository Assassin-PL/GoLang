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
	var j int8 = 27
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
	fmt.Printf("\n")
	fmt.Print(n)
	fmt.Printf("\n")
	fmt.Printf("%v, %T", o, o)
	fmt.Printf("\n")
	fmt.Print(j + int8(k))
	a := 10               //1010
	b := 3                //0011
	fmt.Print(a&b, "\n")  // funkcja boolowska AND
	fmt.Print(a|b, "\n")  // funkcja OR
	fmt.Print(a^b, "\n")  // funkcja XOR
	fmt.Print(a&^b, "\n") // funkcja NAND
	fmt.Print(a<<1, "\n") // przesuniecie bitowe w lewo
	fmt.Print(a>>1, "\n") // przesuniecie bitowe w prawo
	// typy zespolone
	var v complex64 = 1 + 2i //sa to typy float 32 a nie int 32!!!
	var z complex64 = 4 + 7i
	fmt.Printf("%v, %T", v, v)
	fmt.Printf("\n")
	fmt.Print(v+z, "\n") // dodawanie
	fmt.Print(v-z, "\n") // odejmowanie
	fmt.Print(v*z, "\n") // mnozenie
	fmt.Print(v/z, "\n") // dzielenie
	// typy tekstowe string tak jak w c++
	m = "Dupa"
	fmt.Print(actorName + "  " + m + "  ") //konkatynacja jak w pythonie, nie ma typu char bo w stringu sa aliasy do literek a nie chary

}
