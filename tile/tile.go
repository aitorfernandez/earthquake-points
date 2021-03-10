package tile

import (
	"image"
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// Tile ...
type Tile struct{}

// New ...
func New() *Tile {
	return &Tile{}
}

// Draw ...
func (t Tile) Draw() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 256, 256))
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			c := colorful.Hsv(215.0, 1, 1)
			r, g, b := c.RGB255()
			img.SetNRGBA(x, y, color.NRGBA{r, g, b, uint8(100)})
		}
	}
	return img
}
