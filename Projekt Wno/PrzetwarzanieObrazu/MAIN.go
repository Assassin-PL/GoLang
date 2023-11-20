package main

import (
	"flag"
	_ "image/png"
)

func main() {
	//Definicja parametrow:
	var contrast_level float64 = 5.5                 //w %
	var sharpen_level float64 = 5                    //
	var gaussiaNoise_level float64 = stalaszumu * 60 // w %
	var blur_level float64 = stalabluru * 20         // w %
	var distoreted_level float64 = amplituda * 0.02
	var distoreted_freq float64 = pulsacja
	// Definicja flag - uzywam do wywolywania funkcji w przetwarzaniu plikow
	imagePath := flag.String("input", "Przyklad.jpg", "Ścieżka do pliku obrazu")
	sharpen := flag.Bool("sharpen", true, "Wyostrzanie obrazu")
	contrast := flag.Bool("contrast", true, "Uwydatnianie kontrastu")
	gaussiaNoise := flag.Bool("gaussiaNoise", true, "Dodawanie szumu Gausowskiego do obrazu")
	blur := flag.Bool("blur", true, "Funkcja rozmycia Gaussowskiego")
	distoreted := flag.Bool("distoreted", true, "Powoduje znieksztalcennie obrazu")
	changePainting(&contrast_level, &sharpen_level, &gaussiaNoise_level, &blur_level, &distoreted_level, &distoreted_freq, imagePath, sharpen, contrast, gaussiaNoise, blur, distoreted) // funkcja odpowiada za zmiane obrazka
}
