package pikselkapcio

import (
	"image/color"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	//Case 1
	//Test generated image size and code length
	config := Config{
		Scale:            9,
		RandomTextLength: 12,
	}

	codeText, codeImageData := GenerateCode(config)

	if len(codeText) != 12 {
		t.Error("Code length should be 12")
	}

	if codeImageData.Bounds().Size().X != 7*9*12 {
		t.Error("Image width should be 756")
	}

	if codeImageData.Bounds().Size().Y != 7*9 {
		t.Error("Image height should be 63")
	}

	//Case 2
	//test image foreground and background colors for last character
	config = Config{
		CustomWords:        []string{"Testing_is_fun"},
		TextGenerationMode: TextGenerationCustomWords,
		ColorHexStringPairs: []HexStringPair{
			{
				BackgroundColor: "FF0000",
				ForegroundColor: "FFFFFF",
			},
			{
				BackgroundColor: "880088",
				ForegroundColor: "0044FF",
			},
			{
				BackgroundColor: "444444",
				ForegroundColor: "0088CC",
			},
		},
		ColorPairsRotation: ColorPairsRotationSequence,
	}

	codeText, codeImageData = GenerateCode(config)

	//letters of generated code string are always uppercase
	if codeText != "TESTING_IS_FUN" {
		t.Error("Generated code should be 'TESTING_IS_FUN'")
	}

	//if ColorPairsRotationSequence is used, last (14th) charcter will get 2nd out of 3 defined color pairs
	expectedLastCharacterBacgroundColor := color.RGBA{136, 0, 136, 255} //880088
	actualLastCharacterBackgroundColor := codeImageData.At(codeImageData.Bounds().Max.X-1, codeImageData.Bounds().Max.Y-1)

	if actualLastCharacterBackgroundColor != expectedLastCharacterBacgroundColor {
		t.Error("Invalid last characters background color")
	}

	expectedLastCharacterForegroundColor := color.RGBA{0, 68, 255, 255} //0044FF
	actualLastCharacterForeroundColor := codeImageData.At(codeImageData.Bounds().Max.X-1-9, codeImageData.Bounds().Max.Y-1-9)

	if actualLastCharacterForeroundColor != expectedLastCharacterForegroundColor {
		t.Error("Invalid last characters foreground color")
	}

	//Case 3
	//test image width and code case for multibyte utf8 characters
	config.CustomWords = []string{"Zażółć gęślą jaźń"}
	config.Scale = 2

	codeText, codeImageData = GenerateCode(config)

	if codeText != "ZAŻÓŁĆ GĘŚLĄ JAŹŃ" {
		t.Error("Generated code should be 'ZAŻÓŁĆ GĘŚLĄ JAŹŃ'")
	}

	if codeImageData.Bounds().Size().X != 7*2*17 {
		t.Error("Image width should be 238")
	}
}
