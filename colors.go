package pikselkapcio

import (
	"image/color"
	"strconv"
)

//HexStringPair represents pair of hex strings representing background and foreground colors
type HexStringPair struct {
	BackgroundColor string
	ForegroundColor string
}

type colorPair struct {
	backgroundColor color.RGBA
	foregroundColor color.RGBA
}

func generateRGBColorFromHexString(hexString string) color.RGBA {
	decimalColor, err := strconv.ParseUint(hexString, 16, 32)

	if err != nil {
		panic("Invalid hexadecimal color string")
	}

	return color.RGBA{uint8(decimalColor >> 16), uint8((decimalColor >> 8) & 0xFF), uint8(decimalColor & 0xFF), 255}
}

func generateColorPairs(hexStringPairs []HexStringPair) []colorPair {
	colorPairs := []colorPair{}

	for _, hexStringPair := range hexStringPairs {
		colorPairs = append(colorPairs, colorPair{
			backgroundColor: generateRGBColorFromHexString(hexStringPair.BackgroundColor),
			foregroundColor: generateRGBColorFromHexString(hexStringPair.ForegroundColor),
		})
	}

	return colorPairs
}
