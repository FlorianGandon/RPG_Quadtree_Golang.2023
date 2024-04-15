package character

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, width, height int) {

	if !c.Moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) && inpututil.KeyPressDuration(ebiten.KeyControl) == 0 {
			c.Orientation = OrientedRight
			if !blocking[1] {
				c.xInc = 1
				c.Moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && inpututil.KeyPressDuration(ebiten.KeyControl) == 0 {
			c.Orientation = OrientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.Moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) && inpututil.KeyPressDuration(ebiten.KeyControl) == 0 {
			c.Orientation = OrientedUp
			if !blocking[0] {
				c.yInc = -1
				c.Moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) && inpututil.KeyPressDuration(ebiten.KeyControl) == 0 {
			c.Orientation = OrientedDown
			if !blocking[2] {
				c.yInc = 1
				c.Moving = true
			}
		}
	} else {
		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.Moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
			}
		}
	}
	c.XShift = 0
	c.YShift = 0
	switch c.Orientation {
	case OrientedDown:
		c.YShift = c.shift
	case OrientedUp:
		c.YShift = -c.shift
	case OrientedLeft:
		c.XShift = -c.shift
	case OrientedRight:
		c.XShift = c.shift
	}
	if configuration.Global.RoundEarth {
		if c.X <= -1 {
			c.X += width
		}
		if c.Y <= -1 {
			c.Y += height
		}
		if c.X >= width {
			c.X -= width
		}
		if c.Y >= height {
			c.Y -= height
		}
	}
}
