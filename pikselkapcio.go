package pikselkapcio

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//GenerateImageData generates image data
func GenerateImageData() image.PalettedImage {
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewPaletted(rect, []color.Color{color.Black, color.White})
	img.SetColorIndex(50, 50, 1)

	return img
}

//ImageHandler serves generated image as PNG
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := png.Encode(w, GenerateImageData()); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
}

//GenerateRandomText generates pseudo random uppercased string of specified length
func GenerateRandomText(length int) string {
	rand.Seed(time.Now().Unix())

	in := "0123456789abcdefghijklmnopqrstuvwxyz"

	inRune := []rune(in)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return strings.ToUpper(string(inRune)[:length])
}

func generateEmptyPixelMap(textLength int) {
	pixelMap := make([][7]string, textLength*7)
	for x := 0; x < textLength*7; x++ {
		for y := 0; y < 7; y++ {
			pixelMap[x][y] = "000000"
		}

	}
}
