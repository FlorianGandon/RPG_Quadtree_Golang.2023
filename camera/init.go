package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place une caméra.
func (c *Camera) Init(width, height int) {
	if configuration.Global.CameraMode == Static {
		c.FluideX = false
		c.FluideY = false
	}
	if configuration.Global.CameraMode == Normal {
		c.FluideX = false
		c.FluideY = false
	}
	if configuration.Global.CameraMode == Fluide {
		c.FluideX = true
		c.FluideY = true
	}
	if configuration.Global.StopCamera {
		if configuration.Global.NumTileX <= width {
			c.X = configuration.Global.ScreenCenterTileX
		} else {
			c.X = width / 2
		}
		if configuration.Global.NumTileY <= height {
			c.Y = configuration.Global.ScreenCenterTileY
		} else {
			c.Y = height / 2
		}

	}
}
