package character

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw permet d'afficher le personnage dans une *ebiten.Image
// (en pratique, celle qui représente la fenêtre de jeu) en
// fonction des charactéristiques du personnage (position, orientation,
// étape d'animation, etc) et de la position de la caméra (le personnage
// est affiché relativement à la caméra).
// Les arguments sont :
// camXPos et camYPos sont la position de la caméra,
// width et  height int sont la taille du fullContent,
// FluideX et FluideY sont si les mouvement du joueur sont fluide sur l'axe X et Y.
func (c Character) Draw(screen *ebiten.Image, camXPos, camYPos, width, height int, FluideX, FluideY bool) {
	if configuration.Global.RoundEarth {
		// le mode RoundEarth requierd un affichage particulier
		for xTileForDisplay := ((c.X - camXPos) + (configuration.Global.ScreenCenterTileX % width)) % width; xTileForDisplay < configuration.Global.NumTileX; xTileForDisplay += width {
			for yTileForDisplay := ((c.Y - camYPos) + (configuration.Global.ScreenCenterTileY % height)) % height; yTileForDisplay < configuration.Global.NumTileY; yTileForDisplay += height {

				var xPos int
				var yPos int
				if FluideX {
					xPos = (xTileForDisplay) * configuration.Global.TileSize
				} else {
					xPos = (xTileForDisplay)*configuration.Global.TileSize + c.XShift
				}
				if FluideY {
					yPos = (yTileForDisplay) * configuration.Global.TileSize
				} else {
					yPos = (yTileForDisplay)*configuration.Global.TileSize + c.YShift
				}

				if configuration.Global.BetterCharacter {
					c.Shadow.DrawShadow(screen, xPos, yPos)
				}

				if configuration.Global.DebugMode && configuration.Global.BetterCharacter {
					yPos -= 12
				}

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(xPos), float64(yPos))

				shiftX := configuration.Global.TileSize
				if c.Moving && (!configuration.Global.BetterCharacter || !configuration.Global.DebugMode) {
					shiftX += c.animationStep * configuration.Global.TileSize
				}
				shiftY := c.Orientation * configuration.Global.TileSize

				screen.DrawImage(assets.CharacterImage.SubImage(
					image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
				).(*ebiten.Image), op)
			}
		}

	} else {

		// position relatif du character
		xTileForDisplay := c.X - camXPos + configuration.Global.ScreenCenterTileX
		yTileForDisplay := c.Y - camYPos + configuration.Global.ScreenCenterTileY

		var xPos int
		var yPos int

		if FluideX {
			xPos = (xTileForDisplay) * configuration.Global.TileSize
		} else {
			xPos = (xTileForDisplay)*configuration.Global.TileSize + c.XShift
		}
		if FluideY {
			yPos = (yTileForDisplay) * configuration.Global.TileSize
		} else {
			yPos = (yTileForDisplay)*configuration.Global.TileSize + c.YShift
		}

		if configuration.Global.BetterCharacter {
			c.Shadow.DrawShadow(screen, xPos, yPos)
		}

		// le joueur "marche" sur une case donc ses pieds doivent être sur la case
		yPos = yPos - 2

		// effet vole en mode debug
		if configuration.Global.DebugMode && configuration.Global.BetterCharacter {
			yPos -= 12
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(xPos), float64(yPos))
		op.ColorScale.Reset()

		shiftX := configuration.Global.TileSize
		if c.Moving && (!configuration.Global.BetterCharacter || !configuration.Global.DebugMode) {
			shiftX += c.animationStep * configuration.Global.TileSize
		}
		shiftY := c.Orientation * configuration.Global.TileSize
		if configuration.Global.BetterCharacter && configuration.Global.DebugMode {
			op.ColorScale.SetG(float32(configuration.Global.SeedRand.Intn(10)))
			op.ColorScale.SetR(float32(configuration.Global.SeedRand.Intn(10)))
			shiftX = 1 * configuration.Global.TileSize
			shiftY = 4 * configuration.Global.TileSize
		}
		screen.DrawImage(assets.CharacterImage.SubImage(
			image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
		).(*ebiten.Image), op)
	}
}
