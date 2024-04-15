package particle

import "image/color"

type Particle struct {
	Color       color.Color
	Radius      float32
	Orientation float64
	distance    float64
}
