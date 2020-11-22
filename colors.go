package pikselkapcio

import (
	"image/color"
	"strconv"
)

//ColorPair represents pair of background and foreground RGBA colors
type ColorPair struct {
	backgroundColor color.RGBA
	foregroundColor color.RGBA
}

//GenerateRGBColorFromHexString generates valid color from hex string (e.g. "FA8C0D")
func GenerateRGBColorFromHexString(hexString string) color.RGBA {
	decimalColor, err := strconv.ParseUint(hexString, 16, 32)

	if err != nil {
		panic("Invalid hexadecimal color string")
	}

	return color.RGBA{uint8(decimalColor >> 16), uint8((decimalColor >> 8) & 0xFF), uint8(decimalColor & 0xFF), 255}
}

//GetDefaultColorPairs returns default set of background and foreground color pairs
func GetDefaultColorPairs() []ColorPair {
	colorPairs := []ColorPair{}
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("CCCCCC"),
			foregroundColor: GenerateRGBColorFromHexString("888888"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("888888"),
			foregroundColor: GenerateRGBColorFromHexString("CCCCCC"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("00CC00"),
			foregroundColor: GenerateRGBColorFromHexString("97EA97"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("97EA97"),
			foregroundColor: GenerateRGBColorFromHexString("00CC00"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("9797EA"),
			foregroundColor: GenerateRGBColorFromHexString("5C5CDE"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("5C5CDE"),
			foregroundColor: GenerateRGBColorFromHexString("9797EA"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("FF8800"),
			foregroundColor: GenerateRGBColorFromHexString("FFCE97"),
		})
	colorPairs = append(colorPairs,
		ColorPair{
			backgroundColor: GenerateRGBColorFromHexString("FFCE97"),
			foregroundColor: GenerateRGBColorFromHexString("FF8800"),
		})

	return colorPairs
}
