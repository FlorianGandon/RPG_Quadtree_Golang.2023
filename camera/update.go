package camera

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
)

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY int, OrienPlayeur int, width, height int) {

	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case Normal:
		c.updateFollowCharacter(characterPosX, characterPosY, OrienPlayeur, width, height)
	case Fluide:
		c.updateFollowCharacter(characterPosX, characterPosY, OrienPlayeur, width, height)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int, OrienPlayeur int, width, height int) {
	if configuration.Global.StopCamera {
		pointXtheorique := characterPosX - configuration.Global.ScreenCenterTileX
		pointYtheorique := characterPosY - configuration.Global.ScreenCenterTileY
		if pointXtheorique >= 0 && pointXtheorique <= width-configuration.Global.NumTileX {
			c.X = characterPosX
		}
		if ((pointXtheorique == 0 && OrienPlayeur == 2) || (pointXtheorique == width-configuration.Global.NumTileX && OrienPlayeur == 1)) && configuration.Global.CameraMode != 1 {
			c.FluideX = true
		} else if (pointXtheorique == 0 && OrienPlayeur == 1) || (pointXtheorique == width-configuration.Global.NumTileX && OrienPlayeur == 2) {
			c.FluideX = false
		} else if pointXtheorique > 0 && pointXtheorique < width-configuration.Global.NumTileX && configuration.Global.CameraMode != 1 {
			c.FluideX = true
		} else {
			c.FluideX = false
		}

		if pointYtheorique >= 0 && pointYtheorique <= height-configuration.Global.NumTileY {
			c.Y = characterPosY
		}
		if ((pointYtheorique == 0 && OrienPlayeur == 0) || (pointYtheorique == height-configuration.Global.NumTileY && OrienPlayeur == 3)) && configuration.Global.CameraMode != 1 {
			c.FluideY = true
		} else if (pointYtheorique == 0 && OrienPlayeur == 3) || (pointYtheorique == height-configuration.Global.NumTileY && OrienPlayeur == 0) {
			c.FluideY = false
		} else if pointYtheorique > 0 && pointYtheorique < height-configuration.Global.NumTileY && configuration.Global.CameraMode != 1 {
			c.FluideY = true
		} else {
			c.FluideY = false
		}
	} else {
		c.X = characterPosX
		c.Y = characterPosY
	}
}

func (c *Camera) Teleport(newX, newY, width, height int) {
	pointXtheorique := newX - configuration.Global.ScreenCenterTileX
	pointYtheorique := newY - configuration.Global.ScreenCenterTileY
	if configuration.Global.NumTileX > width {
		c.X = width / 2
	} else if pointXtheorique <= -1 {
		c.X = configuration.Global.ScreenCenterTileX
	} else if pointXtheorique+configuration.Global.NumTileX >= width {
		c.X = width - configuration.Global.ScreenCenterTileX
	} else {
		c.X = newX
	}
	if configuration.Global.NumTileY > height {
		c.Y = height / 2
	} else if pointYtheorique <= -1 {
		c.Y = configuration.Global.ScreenCenterTileY
	} else if pointYtheorique+configuration.Global.NumTileY >= height {
		c.Y = height - configuration.Global.ScreenCenterTileY
	} else {
		c.Y = newY
	}

}
