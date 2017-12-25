package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func drawMander(img *image.RGBA) {
	// xMin := -2.5
	xMin := -250
	// yMin := -2.0
	yMin := -200
	// xMax := 2.5
	xMax := 250
	// yMax := 2.0
	yMax := 200

	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			i := 0
			c := complex(float64(x)*0.01, float64(y)*0.01)
			z := complex(0, 0)
			pixelColor := color.RGBA{255, 255, 255, 255}

			for i < 4000 {
				z = cmplx.Pow(z, 2) + c
				if cmplx.Abs(z) > 4 {
					pixelColor = color.RGBA{uint8(i) * 3, uint8(i) * 2, uint8(i), 255}
					break
				}
				i++
			}

			corX := int(x) + 250
			corY := int(y) + 250
			img.Set(corX, corY, pixelColor)
		}
	}
}

func main() {
	// Create an 100 x 50 image
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	drawMander(img)
	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
