https://go.dev/tour/methods/25

package main

import ("golang.org/x/tour/pic"
		"image"
		"image/color"
	   )

type Image struct{}

// Implementations of image.Image interface for Image

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	// The old image generator used 256x256 image
	return image.Rect(0, 0, 256, 256) 
}

func (i Image) At(x, y int) color.Color {
	// Uncomment one of these functions below and
    // comment rest to see the magic (1st one is             
    // default in this case)
	return color.RGBA{uint8(x^y), uint8(x^y), 255, 255}
	//return color.RGBA{(uint8(x)+uint8(y))/2, (uint8(x)+uint8(y))/2, 255, 255}
	//return color.RGBA{uint8(x)*uint8(y), uint8(x)*uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
