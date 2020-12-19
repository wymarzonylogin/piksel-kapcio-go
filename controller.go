package pikselkapcio

import (
	"image/png"
	"log"
	"net/http"
)

//ImageHandler is an example http handler that serves image generated based on provided configuration as PNG file
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	config := Config{
		Scale:              4,
		TextGenerationMode: TextGenerationCustomWords,
		CustomWords:        []string{"Gopher", "Elephant", "Hamster", "Mouse", "Panda"},
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

	codeText, codeImageData := GenerateCode(config)

	//here you should store generated codeText in session

	if err := png.Encode(w, codeImageData); err != nil {
		log.Printf("Error while encoding image for code '%s'", codeText)
	}

	w.Header().Set("Content-Type", "image/png")
}
