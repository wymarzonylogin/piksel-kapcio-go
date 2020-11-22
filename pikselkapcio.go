package pikselkapcio

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"strconv"
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

	rect := image.Rect(0, 0, imageWidth, imageHeight)
	img := image.NewPaletted(rect, []color.Color{color.Black, color.White})
	println("Text: " + text)
	pixelMap := generatePixelMapForText(text)

	for columnIndex, column := range pixelMap {
		for rowIndex, row := range column {
			if row == "FFFFFF" {
				for colOffset := 0; colOffset < config.Scale; colOffset++ {
					for rowOffset := 0; rowOffset < config.Scale; rowOffset++ {
						img.SetColorIndex(columnIndex*config.Scale+colOffset, rowIndex*config.Scale+rowOffset, 1)
					}
				}
			}
		}
	}

	return img
}

//ImageHandler serves generated image as PNG
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := png.Encode(w, GenerateImageData()); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
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

	for characterIndex, character := range text {
		characterMap := GetPaddedCharacterMap(character)

		for lineIndex, line := range characterMap {
			lineBitMap := fmt.Sprintf("%06s", strconv.FormatInt(line, 2)) + "0"

			for bitOffset, bit := range lineBitMap {
				value := "000000"
				if bit == '1' {
					value = "FFFFFF"
				}
				pixelMap[characterIndex*7+bitOffset][lineIndex] = value
			}
		}
	}

	return pixelMap
}
