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
