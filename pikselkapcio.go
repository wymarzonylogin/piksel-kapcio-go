package pikselkapcio

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Config strucure for whole captcha generation
type Config struct {
	Scale int
}

//GenerateImageData generates image data
func GenerateImageData() image.PalettedImage {
	config := Config{Scale: 7}
	text := GenerateRandomText(4)

	imageWidth := 7 * config.Scale * len(text)
	imageHeight := 7 * config.Scale

	rect := image.Rect(0, 0, imageWidth-1, imageHeight-1)
	img := image.NewPaletted(rect, []color.Color{color.Black, color.White})
	println("Text: " + text)
	//pixelMap := generatePixelMapForText(text)
	generatePixelMapForText(text)
	//img.SetColorIndex(20, 20, 1)

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

	alphabet := "0123456789abcdefghijklmnopqrstuvwxyz"

	alphabetRunes := []rune(alphabet)
	rand.Shuffle(len(alphabetRunes), func(i, j int) {
		alphabetRunes[i], alphabetRunes[j] = alphabetRunes[j], alphabetRunes[i]
	})

	return strings.ToUpper(string(alphabetRunes)[:length])
}

func generateEmptyPixelMap(textLength int) [][7]string {
	pixelMap := make([][7]string, textLength*7)
	for x := 0; x < textLength*7; x++ {
		for y := 0; y < 7; y++ {
			pixelMap[x][y] = "000000"
		}

	}

	return pixelMap
}

func generatePixelMapForText(text string) [][7]string {
	pixelMap := generateEmptyPixelMap(len(text))

	for _, character := range text {
		characterMap := GetCharacterMap(character)

		print("Character: ")
		println(character)

		for _, line := range characterMap {
			lineBitMap := fmt.Sprintf("%06s", strconv.FormatInt(line, 2)) + "0"
			println(lineBitMap)
		}
	}

	return pixelMap
}
