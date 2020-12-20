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
		CustomWords:        []string{"Squirrel", "Gopher", "Elephant", "Hamster", "Octopus", "Panda"},
		ColorHexStringPairs: []HexStringPair{
			{
				BackgroundColor: "7FD5EA",
				ForegroundColor: "FFDA87",
			},
			{
				BackgroundColor: "FF9587",
				ForegroundColor: "FFDA87",
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
