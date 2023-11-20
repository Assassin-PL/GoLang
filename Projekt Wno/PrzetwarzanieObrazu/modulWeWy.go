package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

// wczytywanie obrazka
func loadAndProcessImage(imagePath string) image.Image {
	fullPath := filepath.Join("Obrazy", imagePath)
	file, err := os.Open(fullPath) // funkcja zwrocila blad nie mozna wykonac, sprawdzam tu poprawnosc zapisu pliku
	if err != nil {
		fmt.Println("Błąd podczas otwierania pliku obrazu:", err)
		os.Exit(1) // uzylem tej funkcji dla tego ze jesli bedzie napotkany blad to spowoduje natychmiastowe zamkniecie systemu zeby nie uszkodzic obrazu zrodlowego
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Błąd podczas dekodowania pliku obrazu:", err)
		os.Exit(1)
	}

	return img // funkcja nie zwraca kopii tylko operuje na obrazie wiec trzeba uwazac przy kazdej zmianie tej zmiennej bo ma wplywa na orginal
}

// blok zapisu do pliku
func saveImage(img image.Image, outputPath string) {
	fullPath := filepath.Join("Przetworzone Obrazy", outputPath)
	file, err := os.Create(fullPath)
	if err != nil { // funkcja zwrocila blad nie mozna wykonac, sprawdzam tu poprawnosc zapisu pliku
		fmt.Println("Błąd podczas tworzenia pliku wyjściowego:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Zapisz obraz w formacie JPEG
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		fmt.Println("Błąd podczas zapisywania obrazu:", err)
		os.Exit(1)
	}
}
