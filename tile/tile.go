package tile

import (
	"image"
	"image/color"

	"github.com/aitorfernandez/earthquake-points/quake"
	"github.com/lucasb-eyer/go-colorful"
)

// Tile ...
type Tile struct{}

// New ...
func New() *Tile {
	return &Tile{}
}

// Draw ...
func (t Tile) Draw(a, b int) image.Image {
	q := quake.New()

	size := 256
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		yy := y + (256 * b)
		if yy > q.Loc.Y-10 && yy < q.Loc.Y+10 {
			for x := 0; x < size; x++ {
				xx := x + (256 * a)
				if xx > q.Loc.X-10 && xx < q.Loc.X+10 {
					c := colorful.Hsv(215.0, 1, 1)
					r, g, b := c.RGB255()
					img.SetNRGBA(x, y, color.NRGBA{r, g, b, uint8(100)})
				}
			}
		}
	}
	return img
}
