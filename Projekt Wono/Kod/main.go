package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Zmienna globalna reprezentująca nazwisko postaci.
var czlek = "John Deep"

const (
	_ = iota
	// a //
)

// Struktura reprezentująca postać lekarza w serialu Doctor Who.
type Doctor struct {
	number    int
	actorName string
	companion []string
}

// Struktura Animal reprezentuje zwierzę z informacjami o nazwie i pochodzeniu.
type Animal struct {
	Name   string `required:"true" max:"100"` // dodawanie Tagu do struktury , taka informacja gdy deklarujemy strukture
	Origin string
}

// Struktura Bird dziedziczy z Animal i dodaje dodatkowe informacje o prędkości i zdolności do latania.
type Bird struct {
	Animal   // Oszukane dziedziczenie struktury Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// Przestrzeń deklaracyjna

	// Deklaracja stałej lokalnej.
	const MY_CONST int = 1
	// Shadowing zmiennej 'a' zadeklarowanej na poziomie pakietu w funkcji main.
	const a int16 = 69

	// Deklaracja tablicy o stałej wielkości.
	grades := [...]int32{97, 85, 93}

	var students [3]string // Deklaracja tablicy z imionami studentów.

	// Deklaracja macierzy jednostkowej.
	var identityMatrix [3][3]int = [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	// Deklaracja struktury reprezentującej postać lekarza w Doctor Who.
	aDoctor := Doctor{
		number:    3,
		actorName: "Janusz Korwin Mikke",
		companion: []string{
			"Grzegorz Braun",
			"Krzysiu Bosak",
			"Konrad Berkowicz",
		},
	}

	// Deklaracja mapy przechowującej populację miast w Polsce.
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

	GRADES := &grades // Kopiowanie tablicy do tablicy przez referencję.

	// Inicjalizacja slices.
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := b[:]
	d := b[3:]
	e := b[:6]
	f := b[3:6]

	g := make([]int, 3, 100) // Trzeci argument określa maksymalną pojemność slices, a jej rozmiaru nie można później zmieniać.

	// Inicjalizacja obiektu Bird (ptaszek).
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	ptaszek := Bird{
		Animal:   Animal{Name: "Twoja stara", Origin: "dupa"},
		SpeedKPH: 69,
		CanFly:   true,
	}
	// Modyfikacja pól struktury ptaszek.
	ptaszek.Name = "Kruk"
	ptaszek.Origin = "gothic II NK"
	ptaszek.SpeedKPH = 69
	ptaszek.CanFly = false

	students[0] = "Lisa" // Przypisanie wartości do elementu tablicy.

	i := 42
	var n bool = true
	var j int8 = 27
	k := 99
	var m string = "NULL"
	o := 1 == 2         // Inicjalizacja zmiennej boolowskiej.
	m = strconv.Itoa(k) // Konwersja liczby całkowitej na string.

	// Funkcje na tablicach.
	g = append(g, 1)                                                                               // Dodawanie elementu do slices.
	g = append(g, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}...) // Dodawanie innej tablicy do slices.
	g = g[1:]                                                                                      // Usuwanie pierwszego elementu slices.
	g = g[:len(g)-1]                                                                               // Usuwanie ostatniego elementu slices.
	g = append(g[:len(g)/2], g[len(g)/2+1:]...)                                                    // Wycinanie elementu środkowego slices.

	// Funkcje na mapach.
	statePopulations["Zbąszynek"] = 69
	delete(statePopulations, "Zbąszynek")
	_, ok := statePopulations["Zbąszynek"] // Sprawdzanie, czy istnieje klucz w mapie.

	// Wyświetlanie wyników.
	fmt.Println(field.Tag)
	fmt.Print("\n")
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
	fmt.Printf("Studenci: %v", students)
	fmt.Print("\n")
	fmt.Printf("Oceny: %v", grades)
	fmt.Print("\n")
	fmt.Printf("Długość ocen: %v", len(GRADES))
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

	// Typy zespolone.
	var v complex64 = 1 + 2i
	var z complex64 = 4 + 7i
	printFormattedValues(v)
	fmt.Print(v+z, "\n") // Dodawanie
	fmt.Print(v-z, "\n") // Odejmowanie
	fmt.Print(v*z, "\n") // Mnożenie
	fmt.Print(v/z, "\n") // Dzielenie

	// Typy tekstowe (string) takie jak w języku C++.
	m = "Dupa"
	fmt.Print(czlek + "  " + m + "  ") // Konkatenacja, nie ma typu 'char', bo w stringu są aliasy do liter a nie chary
	fmt.Print("\n")
	fmt.Println(aDoctor.actorName)
	fmt.Print("\n")
}

// Funkcja do wyświetlania wartości.
func printValues(values ...interface{}) {
	for _, v := range values {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

// Funkcja do wyświetlania sformatowanych wartości.
func printFormattedValues(values ...interface{}) {
	for _, v := range values {
		fmt.Printf("%v, %T\n", v, v)
	}
}

// Funkcja do wyświetlania operacji bitowych.
func printBitwiseOperations(a, b int) {
	fmt.Print(a&b, "\n")  // AND
	fmt.Print(a|b, "\n")  // OR
	fmt.Print(a^b, "\n")  // XOR
	fmt.Print(a&^b, "\n") // NAND
	fmt.Print(a<<1, "\n") // Przesunięcie bitowe w lewo
	fmt.Print(a>>1, "\n") // Przesunięcie bitowe w prawo
}
