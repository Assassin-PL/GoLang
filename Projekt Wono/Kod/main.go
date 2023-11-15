package main

import (
	"fmt"
	"strconv"
)

var czlek = "John Deep"

const (
	_ = iota
	//a //
)

// stworzenie typu struct
type Doctor struct {
	number    int
	actorName string
	companion []string
}

func main() {
	//Przestrzen deklaracyjne
	const MY_CONST int = 1
	const a int16 = 69 // shadowing nastopilo "przykrycie zmiennej a zapisanej w funkcji"

	grades := [...]int32{ //Tablica 1D
		97, 85, 93,
	}

	var studnets [3]string

	var identityMatrix [3][3]int = [3][3]int{
		{
			1, 0, 0,
		},
		{
			0, 1, 0,
		},
		{
			0, 0, 1,
		},
	}
	//deklaracja struktury
	aDoctor := Doctor{
		number:    3,
		actorName: "Janusz Korwin Mikke",
		companion: []string{
			"Grzegorz Braun",
			"Krzysiu Bosak",
			"Konrad Berkowicz",
		},
	}

	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal   //takie oszukane dziedziczenie
		SpeedKPH float32
		CanFly   bool
	}

	// deklaracje map (odpowiednik slownikow w jezyku python ,  gdzie z lewej jest klucz a z prawej wartosc)

	//statePopulations := make(map[string]int) // deklaracja may
	statePopulations := map[string]int{
		"Warszawa":    1793579,
		"Kraków":      769498,
		"Łódź":        682679,
		"Wrocław":     641607,
		"Poznań":      538633,
		"Gdańsk":      470907,
		"Szczecin":    401907,
		"Bydgoszcz":   350178,
		"Lublin":      339784,
		"Białystok":   294337,
		"Katowice":    291422,
		"Gdynia":      245454,
		"Częstochowa": 223152,
		"Radom":       213029,
		"Kielce":      200635,
		"Gliwice":     185450,
		"Toruń":       202562,
		"Bytom":       170537,
		"Zabrze":      176382,
		// Dodaj więcej miast w Polsce w razie potrzeby.
	}

	GRADES := &grades // kopjowanie tablic do tablicy przez referencje, czyli dajemy alias a  nie kopjujemy calej tablicy

	//slices tak jak w pythonie chcemy sobie podzilic tablice

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := b[:]
	d := b[3:]
	e := b[:6]
	f := b[3:6]

	g := make([]int, 3, 100) // 3 agrument pokazuje jak duzo moze byc maksymalnie ta tablica pozniej jej rozmiaru nie mozna zmieniac

	// smietnik zmiennyuch
	ptaszek := Bird{
		Animal:   Animal{Name: "Twoja stara", Origin: "dupa"},
		SpeedKPH: 69,
		CanFly:   true,
	}
	ptaszek.Name = "Kruk"
	ptaszek.Origin = "gothic II NK"
	ptaszek.SpeedKPH = 69
	ptaszek.CanFly = false
	studnets[0] = "Lisa"
	i := 42
	var n bool = true
	var j int8 = 27
	k := 99
	var m string = "NULL"
	o := 1 == 2 // inicjalizacja zmiennych boolowskich
	m = strconv.Itoa(k)

	//funkcje na tablicach
	g = append(g, 1)                                                                               // dodowanie zmiennej do tablicy
	g = append(g, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}...) // ciekamy spob na dodowanie tablicy do tablicy tablica 1D staje sie tablica 2D
	g = g[1:]                                                                                      // odejmowanie pierwszego elementuj
	g = g[:len(g)-1]                                                                               //odejmowanie ostatniego elementu
	g = append(g[:len(g)/2], g[len(g)/2+1:]...)                                                    //wycinanie wyrazu srodkowego

	//funkcje na mapach
	statePopulations["Zbąszynek"] = 69
	delete(statePopulations, "Zbąszynek")
	_, ok := statePopulations["Zbąszynek"] // zwraca true jesli istnieje taki klucz
	//funkcje wyswietlania
	fmt.Println(statePopulations)
	fmt.Print("\n")
	fmt.Println(ok)
	fmt.Print("\n")
	fmt.Println(b, c, d, e, f)
	fmt.Print("\n")
	fmt.Println(g)
	fmt.Print("\n")
	fmt.Printf("Macierz jednostkowa: %v", identityMatrix)
	fmt.Print("\n")
	fmt.Printf("Students: %v", studnets)
	fmt.Print("\n")
	fmt.Printf("grades: %v", grades)
	fmt.Print("\n")
	fmt.Printf("grades dlugosc: %v", len(GRADES))
	fmt.Print("\n")
	printValues(a)
	printValues(MY_CONST)
	printValues(i, j, k)
	printFormattedValues(i, k)
	fmt.Print(czlek + "\n")
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
	fmt.Print(czlek + "  " + m + "  ") // konkatenacja jak w Pythonie, nie ma typu char, bo w stringu są aliasy do liter a nie chary
	fmt.Print("\n")
	fmt.Println(aDoctor.actorName)
	fmt.Print("\n")
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
