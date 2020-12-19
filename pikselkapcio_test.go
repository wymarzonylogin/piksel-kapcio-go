package pikselkapcio

import (
	"testing"
)

func TestGenerateCode(t *testing.T) {
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

	config = Config{
		CustomWords:        []string{"Testing_is_fun"},
		TextGenerationMode: TextGenerationCustomWords,
	}

	codeText, _ = GenerateCode(config)

	//letters of generated code string are always uppercase
	if codeText != "TESTING_IS_FUN" {
		t.Error("Generated code should be 'TESTING_IS_FUN'")
	}
}
