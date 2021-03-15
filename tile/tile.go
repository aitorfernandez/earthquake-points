package tile

import (
	"image"
	"image/color"

	"github.com/aitorfernandez/earthquake-points/quake"
	"github.com/lucasb-eyer/go-colorful"
)

// Tile structs ...
type Tile struct {
	Image *image.NRGBA
	Size  int
	X     int
	Y     int
}

// New returns a new initialize Tile struct.
func New(x, y int) *Tile {
	size := 256
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return &Tile{
		Image: img,
		Size:  size,
		X:     x,
		Y:     y,
	}
}

// GenerateImage ...
func (t *Tile) GenerateImage(q *quake.Quake) {
	for y := 0; y < t.Size; y++ {
		yy := y + (256 * t.Y)
		if yy > q.Loc.Y-10 && yy < q.Loc.Y+10 {
			for x := 0; x < t.Size; x++ {
				xx := x + (256 * t.X)
				if xx > q.Loc.X-10 && xx < q.Loc.X+10 {
					c := colorful.Hsv(215.0, 1, 1)
					r, g, b := c.RGB255()
					t.Image.SetNRGBA(x, y, color.NRGBA{r, g, b, uint8(100)})
				}
			}
		}
	}
}
