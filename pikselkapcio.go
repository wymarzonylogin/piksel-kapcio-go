package pikselkapcio

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

//GenerateCode generates image data, returns both code as string and image data
func GenerateCode(customConfig Config) (string, image.Image) {
	config := mergeConfig(customConfig)
	text := getText(config)
	colorPairs := generateColorPairs(config.ColorHexStringPairs)
	rect := image.Rect(0, 0, 7*config.Scale*len(text), 7*config.Scale)
	img := image.NewRGBA(rect)
	pixelMap := generatePixelColorMapForText(text, colorPairs)

	for columnIndex, column := range pixelMap {
		for rowIndex, colorValue := range column {
			for colOffset := 0; colOffset < config.Scale; colOffset++ {
				for rowOffset := 0; rowOffset < config.Scale; rowOffset++ {
					img.SetRGBA(columnIndex*config.Scale+colOffset, rowIndex*config.Scale+rowOffset, colorValue)
				}
			}
		}
	}

	return text, img
}

func generatePixelColorMapForText(text string, colorPairs []colorPair) [][7]color.RGBA {
	pixelMap := make([][7]color.RGBA, len(text)*7)
	rand.Seed(time.Now().UnixNano())

	for characterIndex, character := range text {
		characterMap := getPaddedCharacterMap(character)
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
