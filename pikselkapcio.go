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
	"time"
)

//Config strucure for whole captcha generation
type Config struct {
	Scale int
}

func getDefaultConfig() Config {
	return Config{Scale: 7}
}

//GenerateImageData generates image data
func GenerateImageData(config Config) image.Image {
	text := GenerateRandomText(4)
	rect := image.Rect(0, 0, 7*config.Scale*len(text), 7*config.Scale)
	img := image.NewRGBA(rect)
	pixelMap := generatePixelColorMapForText(text)

	for columnIndex, column := range pixelMap {
		for rowIndex, colorValue := range column {

			for colOffset := 0; colOffset < config.Scale; colOffset++ {
				for rowOffset := 0; rowOffset < config.Scale; rowOffset++ {
					img.SetRGBA(columnIndex*config.Scale+colOffset, rowIndex*config.Scale+rowOffset, colorValue)
				}
			}
		}
	}

	return img
}

//ImageHandler serves generated image as PNG
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	config := Config{
		Scale: 22,
	}

	if err := png.Encode(w, GenerateImageData(config)); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
}

func generateEmptyPixelMap(textLength int) [][7]color.RGBA {
	pixelMap := make([][7]color.RGBA, textLength*7)
	for x := 0; x < textLength*7; x++ {
		for y := 0; y < 7; y++ {
			pixelMap[x][y] = color.RGBA{}
		}

	}

	return pixelMap
}

func generatePixelColorMapForText(text string) [][7]color.RGBA {
	pixelMap := generateEmptyPixelMap(len(text))
	rand.Seed(time.Now().UnixNano())
	colorPairs := GetDefaultColorPairs()

	for characterIndex, character := range text {
		characterMap := GetPaddedCharacterMap(character)
		colorPair := colorPairs[rand.Intn(len(colorPairs))]

		for lineIndex, line := range characterMap {
			lineBitMap := fmt.Sprintf("%06s", strconv.FormatInt(line, 2)) + "0"

			for bitOffset, bit := range lineBitMap {
				value := color.RGBA{}
				if bit == '1' {
					value = colorPair.foregroundColor
				} else {
					value = colorPair.backgroundColor
				}
				pixelMap[characterIndex*7+bitOffset][lineIndex] = value
			}
		}
	}

	return pixelMap
}
