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
	pixelMap := generatePixelColorMapForText(text, colorPairs, config.ColorPairsRotation)

	//number of columns: (length of generated code string) * 7 (which is result image width/scale)
	//number of rows: 7 (which is result image height/scale)
	for columnIndex, column := range pixelMap {
		for rowIndex, colorValue := range column {
			fillScaledUpPixel(columnIndex, rowIndex, config.Scale, colorValue, *img)
		}
	}

	return text, img
}

func fillScaledUpPixel(columnIndex, rowIndex, scale int, colorValue color.RGBA, img image.RGBA) {
	for colOffset := 0; colOffset < scale; colOffset++ {
		for rowOffset := 0; rowOffset < scale; rowOffset++ {
			img.SetRGBA(columnIndex*scale+colOffset, rowIndex*scale+rowOffset, colorValue)
		}
	}
}

func generatePixelColorMapForText(text string, colorPairs []colorPair, colorPairsRotation int8) [][7]color.RGBA {
	pixelMap := make([][7]color.RGBA, len(text)*7)
	rand.Seed(time.Now().UnixNano())

	for characterIndex, character := range text {
		characterMap := getPaddedCharacterMap(character)
		var colorPair colorPair

		if colorPairsRotation == ColorPairsRotationSequence {
			colorPair = colorPairs[characterIndex%len(colorPairs)]
		} else {
			colorPair = colorPairs[rand.Intn(len(colorPairs))]
		}

		for lineIndex, line := range characterMap {
			lineBitMap := fmt.Sprintf("%06s", strconv.FormatInt(int64(line), 2)) + "0"

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
