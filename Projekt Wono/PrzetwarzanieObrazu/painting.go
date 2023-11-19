package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"os"
	"path/filepath"
)

func changePainting(
	contrast_level *float64, sharpen_level *float64, gaussiaNoise_level *float64, blur_level *float64, distoreted_level *float64, distoreted_freq *float64,
	imagePath *string, sharpen *bool, contrast *bool, gaussiaNoise *bool, blur *bool, distoreted *bool,
) {
	// zmiana ustawien flag
	*contrast = !*contrast
	*sharpen = !*sharpen
	*gaussiaNoise = !*gaussiaNoise
	*gaussiaNoise = !*gaussiaNoise
	*blur = !*blur
	*distoreted = !*distoreted

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
	if *gaussiaNoise {
		image = applyGaussianNoise(image, *gaussiaNoise_level)
	}
	if *blur {
		image = applyGaussianBlur(image, *blur_level)
	}
	if *sharpen {
		image = applySharpening(image, *sharpen_level)
	}
	if *contrast {
		image = applyContrastEnhancement(image, *contrast_level)
	}
	if *distoreted {
		image = applyDistortion(image, *distoreted_level, *distoreted_freq)
	}

	// Zapisz przetworzony obraz w folderze "Przetworzone Obrazy"
	outputPath := filepath.Join("NoweObrazy", "output.jpg")
	saveImage(image, outputPath)
}
