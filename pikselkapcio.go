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

const (
	//TextGenerationRandom should be used for pseudorandom alphanumeric string for code (it is a default value)
	TextGenerationRandom = 1
	//TextGenerationCustomWords should be used to generate code from provided custom words list
	TextGenerationCustomWords = 2
)

//Config strucure for whole captcha generation
type Config struct {
	Scale               int
	TextGenerationMode  int
	RandomTextLength    int
	CustomWords         []string
	ColorHexStringPairs []HexStringPair
}

func getDefaultConfig() Config {
	return Config{
		Scale:               5,
		TextGenerationMode:  TextGenerationRandom,
		RandomTextLength:    4,
		CustomWords:         getDefaultCustomWords(),
		ColorHexStringPairs: getDefaultHexStringPairs(),
	}
}

//GenerateImageData generates image data
func GenerateImageData(customConfig Config) image.Image {
	config := mergeConfig(customConfig)
	text := GetText(config)
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

	return img
}

func generatePixelColorMapForText(text string, colorPairs []colorPair) [][7]color.RGBA {
	pixelMap := make([][7]color.RGBA, len(text)*7)
	rand.Seed(time.Now().UnixNano())

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

//ImageHandler is an example http handler that serves image generated based on provided configuration as PNG file
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	config := Config{
		Scale:              4,
		RandomTextLength:   36,
		TextGenerationMode: TextGenerationCustomWords,
		CustomWords:        []string{"wYmaRzony", "loGIN", "Smacznego"},
		ColorHexStringPairs: []HexStringPair{
			{
				BackgroundColor: "FF0000",
				ForegroundColor: "FFFFFF",
			},
			{
				BackgroundColor: "FFFFFF",
				ForegroundColor: "00FF00",
			},
		},
	}

	if err := png.Encode(w, GenerateImageData(config)); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
}

func mergeConfig(customConfig Config) Config {
	config := getDefaultConfig()

	if customConfig.Scale != 0 {
		config.Scale = customConfig.Scale
	}

	if customConfig.TextGenerationMode != 0 {
		config.TextGenerationMode = customConfig.TextGenerationMode
	}

	if customConfig.RandomTextLength > 0 && customConfig.RandomTextLength < 37 {
		config.RandomTextLength = customConfig.RandomTextLength
	}

	if len(customConfig.CustomWords) > 0 {
		config.CustomWords = customConfig.CustomWords
	}

	if len(customConfig.ColorHexStringPairs) > 0 {
		config.ColorHexStringPairs = customConfig.ColorHexStringPairs
	}

	return config
}
