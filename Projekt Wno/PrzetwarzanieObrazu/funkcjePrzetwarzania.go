package main

import (
	"image"
	"image/color"
	_ "image/png"
	"math"
	"math/rand"

	"github.com/disintegration/imaging"
)

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

func applyGaussianBlur(img image.Image, sigma float64) image.Image {
	// Zastosuj rozmycie gaussowskie
	blurredImage := imaging.Blur(img, sigma)

	return blurredImage
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

func applyDistortion(img image.Image, amplitude float64, frequency float64) image.Image {
	bounds := img.Bounds()
	distortedImage := imaging.Clone(img) // Tworzymy kopię obrazu, aby nie zmieniać oryginału

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Oblicz zniekształcenie dla współrzędnych x i y za pomocą funkcji sinusoidalnej
			distortedX := int(float64(x) + amplitude*math.Sin(2.0*math.Pi*frequency*float64(y)/float64(bounds.Max.Y)))
			distortedY := int(float64(y) + amplitude*math.Sin(2.0*math.Pi*frequency*float64(x)/float64(bounds.Max.X)))

			// Sprawdzamy, czy nowe współrzędne mieszczą się w granicach obrazu
			if distortedX >= bounds.Min.X && distortedX < bounds.Max.X && distortedY >= bounds.Min.Y && distortedY < bounds.Max.Y {
				// Pobieramy kolor z przekształconych współrzędnych
				distortedColor := img.At(distortedX, distortedY)
				distortedColorRGB := color.RGBAModel.Convert(distortedColor).(color.RGBA)
				// Ustawiamy nowy kolor w kopii obrazu
				distortedImage.Set(x, y, distortedColorRGB)
			}
		}
	}

	return distortedImage
}
