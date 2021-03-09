package quake

import "github.com/aitorfernandez/earthquake-points/math"

// Quake ...
type Quake struct {
	ID   string
	Lat  float64
	Long float64
	Loc  math.Vec2
}

// New ...
func New() *Quake {
	return &Quake{
		ID:   "ci39575519",
		Lat:  35.8715,
		Long: -117.679,
	}
}
