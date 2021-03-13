package quake

import (
	"math"

	m "github.com/aitorfernandez/earthquake-points/math"
)

// Quake ...
type Quake struct {
	ID  string
	Lat float64
	Lon float64
	Loc m.Vec2
}

func degreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func latLonToOffsets(lat, long float64) m.Vec2 {
	fe := 180
	r := 1024 / (2 * math.Pi)

	latRad := degreesToRadians(lat)
	lonRad := degreesToRadians(long + float64(fe))

	x := int(math.Floor(lonRad * r))

	yEqu := r * math.Log(math.Tan(math.Pi/4+latRad/2))
	y := int(math.Floor(1024/2 - yEqu))

	return m.Vec2{X: x, Y: y}
}

// New ...
func New() *Quake {
	q := &Quake{
		ID: "ci39575519",
		// Lat: 35.8715,
		// Lon: -117.679,
		Lat: 41.145556,
		Lon: -73.995,
	}
	q.Loc = latLonToOffsets(q.Lat, q.Lon)
	return q
}
