package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func main() {
	// Definicja flag
	imagePath := flag.String("input", "Przyklad.jpg", "Ścieżka do pliku obrazu")
	//imagePath := filepath.Join(currentDir, "Obrazy", "Przyklad.jpg")
	sharpen := flag.Bool("sharpen", false, "Wyostrzanie obrazu")
	contrast := flag.Bool("contrast", false, "Uwydatnianie kontrastu")
	// Dodaj inne flagi dla innych efektów/filtrów

	// Parsowanie flag
	flag.Parse()

	// Sprawdzenie, czy podano ścieżkę do pliku obrazu
	if *imagePath == "" {
		fmt.Println("Podaj ścieżkę do pliku obrazu.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Wczytaj obraz
	image := loadAndProcessImage(*imagePath)

	// Zastosuj filtry/efekty
	if *sharpen {
		image = applySharpening(image)
	}
	if *contrast {
		image = applyContrastEnhancement(image)
	}
	// Dodaj inne warunki dla innych efektów/filtrów

	// Zapisz przetworzony obraz w folderze "Przetworzone Obrazy"
	outputPath := filepath.Join("NoweObrazy", "output.jpg")
	saveImage(image, outputPath)
}

func loadAndProcessImage(imagePath string) image.Image {
	fullPath := filepath.Join("Obrazy", imagePath)
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("Błąd podczas otwierania pliku obrazu:", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Błąd podczas dekodowania pliku obrazu:", err)
		os.Exit(1)
	}

	return img
}

func applySharpening(img image.Image) image.Image {
	// Implementacja wyostrzania obrazu
	// ...
	return img
}

func applyContrastEnhancement(img image.Image) image.Image {
	// Implementacja uwydatniania kontrastu
	// ...
	return img
}

func saveImage(img image.Image, outputPath string) {
	fullPath := filepath.Join("Przetworzone Obrazy", outputPath)
	file, err := os.Create(fullPath)
	if err != nil {
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
