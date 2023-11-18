package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func main() {
	//Definicja parametrow:
	var contrast_level float64 = 0.5     //w %
	var sharpen_level float64 = 1        //
	var gaussiaNoise_level float64 = 0.1 //
	// Definicja flag
	imagePath := flag.String("input", "Przyklad.jpg", "Ścieżka do pliku obrazu")
	sharpen := flag.Bool("sharpen", true, "Wyostrzanie obrazu")
	contrast := flag.Bool("contrast", true, "Uwydatnianie kontrastu")
	gaussiaNoise := flag.Bool("gaussiaNoise", true, "Dodawanie szumu Gausowskiego do obrazu")
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
	if *gaussiaNoise {
		image = applyGaussianNoise(image, gaussiaNoise_level)
	}
	if *sharpen {
		image = applySharpening(image, sharpen_level)
	}
	if *contrast {
		image = applyContrastEnhancement(image, contrast_level)
	}
	// Dodaj inne warunki dla innych efektów/filtrów

	// Zapisz przetworzony obraz w folderze "Przetworzone Obrazy"
	outputPath := filepath.Join("NoweObrazy", "output.jpg")
	saveImage(image, outputPath)
}

// wczytywanie obrazka
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

// bloki funkcji filtrow
func applySharpening(img image.Image, level float64) image.Image {
	// Przykładowe wyostrzanie obrazu o 10%
	sharpenedImage := imaging.Sharpen(img, level)

	return sharpenedImage
}

func applyContrastEnhancement(img image.Image, level float64) image.Image {
	// Implementacja uwydatniania kontrastu
	contrastedImage := imaging.AdjustContrast(img, level)

	return contrastedImage
}

func applyGaussianNoise(img image.Image, level float64) image.Image {
	bounds := img.Bounds()
	noisyImage := imaging.Clone(img) // Tworzymy kopię obrazu, aby nie zmieniać oryginału

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)

			// Konwersja do koloru RGBA (który jest implementacją interfejsu color.Color)
			originalRGBA := color.RGBAModel.Convert(originalColor).(color.RGBA)

			// Dodajemy szum gaussowski do każdego kanału koloru (R, G, B)
			noisyR := addGaussianNoise(float64(originalRGBA.R), level)
			noisyG := addGaussianNoise(float64(originalRGBA.G), level)
			noisyB := addGaussianNoise(float64(originalRGBA.B), level)

			// Ustawiamy nowy kolor z dodanym szumem
			noisyColor := color.RGBA{
				R: uint8(noisyR),
				G: uint8(noisyG),
				B: uint8(noisyB),
				A: originalRGBA.A,
			}

			// Ustawiamy nowy kolor w kopii obrazu
			noisyImage.Set(x, y, noisyColor)
		}
	}

	return noisyImage
}

func addGaussianNoise(value float64, level float64) float64 {
	// Dodajemy szum gaussowski do wartości koloru
	noise := rand.NormFloat64() * level
	noisyValue := value + noise

	// Ograniczamy wartość do zakresu 0-255
	if noisyValue < 0 {
		noisyValue = 0
	} else if noisyValue > 255 {
		noisyValue = 255
	}

	return noisyValue
}

// blok zapisu do pliku
func saveImage(img image.Image, outputPath string) {
	fullPath := filepath.Join("Przetworzone Obrazy", outputPath)
	file, err := os.Create(fullPath)
	if err != nil { // funkcja zwrocila blad nie mozna wykonac
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
