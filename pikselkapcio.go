package pikselkapcio

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
)

//GenerateImageData generates image data
func GenerateImageData() image.PalettedImage {
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewPaletted(rect, []color.Color{color.Black, color.White})
	img.SetColorIndex(50, 50, 1)

	return img
}

//ImageHandler serves generated image as PNG
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := png.Encode(w, GenerateImageData()); err != nil {
		log.Println("Error while encoding image.")
	}

	w.Header().Set("Content-Type", "image/png")
}
