package portal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particle"
	"image"
)

func (p *Portal) Draw(screen *ebiten.Image, camXPos, camYPos int, Xshift, Yshift int, width, height int, FluideX, FluideY bool) {
	// on fait la même chose pour les deux téléporteur
	for _, teleport := range []*teleport{p.teleportA, p.teleportB} {

		if configuration.Global.RoundEarth {
			if teleport.exist {
				for relatifX := ((teleport.X - camXPos) + (configuration.Global.ScreenCenterTileX % width)) % width; relatifX <= configuration.Global.NumTileX; relatifX += width {
					for relatifY := ((teleport.Y - camYPos) + (configuration.Global.ScreenCenterTileY % height)) % height; relatifY <= configuration.Global.NumTileY; relatifY += height {

						// on affiche si le téléporteur est dans l'écran
						var condition_affichage bool
						if configuration.Global.CameraMode == 2 {
							condition_affichage = relatifX >= -1 && relatifY >= -1 && relatifX <= configuration.Global.NumTileX && relatifY <= configuration.Global.NumTileY
						} else {
							condition_affichage = relatifX >= 0 && relatifY >= 0 && relatifX < configuration.Global.NumTileX && relatifY < configuration.Global.NumTileY

						}
						if condition_affichage {
							// on place l'image
							op := &ebiten.DrawImageOptions{}
							var posX float64
							var posY float64

							if FluideX {
								posX = float64((relatifX)*configuration.Global.TileSize - Xshift)
							} else {
								posX = float64((relatifX) * configuration.Global.TileSize)
							}

							if FluideY {
								posY = float64((relatifY)*configuration.Global.TileSize - Yshift)
							} else {
								posY = float64((relatifY) * configuration.Global.TileSize)
							}
							op.GeoM.Translate(posX, posY)

							// on découpe l'image pour que le portail disparaisse progressivement
							if (relatifX*configuration.Global.TileSize-Xshift) > (configuration.Global.NumTileX-1)*configuration.Global.TileSize && (relatifY*configuration.Global.TileSize-Yshift) > (configuration.Global.NumTileY-1)*configuration.Global.TileSize && configuration.Global.CameraMode == 2 {
								screen.DrawImage(assets.FloorImage.SubImage(
									image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize-((relatifX*configuration.Global.TileSize-Xshift)-(configuration.Global.NumTileX-1)*configuration.Global.TileSize), teleport.topleftImageY+configuration.Global.TileSize-((relatifY*configuration.Global.TileSize-Yshift)-(configuration.Global.NumTileY-1)*configuration.Global.TileSize)),
								).(*ebiten.Image), op)
							} else if (relatifX*configuration.Global.TileSize-Xshift) > (configuration.Global.NumTileX-1)*configuration.Global.TileSize && configuration.Global.CameraMode == 2 {
								screen.DrawImage(assets.FloorImage.SubImage(
									image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize-((relatifX*configuration.Global.TileSize-Xshift)-(configuration.Global.NumTileX-1)*configuration.Global.TileSize), teleport.topleftImageY+configuration.Global.TileSize),
								).(*ebiten.Image), op)
							} else if (relatifY*configuration.Global.TileSize-Yshift) > (configuration.Global.NumTileY-1)*configuration.Global.TileSize && configuration.Global.CameraMode == 2 && (relatifY*configuration.Global.TileSize-Yshift) <= (configuration.Global.NumTileY)*configuration.Global.TileSize {
								screen.DrawImage(assets.FloorImage.SubImage(
									image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize, teleport.topleftImageY+configuration.Global.TileSize-((relatifY*configuration.Global.TileSize-Yshift)-(configuration.Global.NumTileY-1)*configuration.Global.TileSize)),
								).(*ebiten.Image), op)
							} else if (relatifY*configuration.Global.TileSize - Yshift) <= (configuration.Global.NumTileY)*configuration.Global.TileSize {
								screen.DrawImage(assets.FloorImage.SubImage(
									image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize, teleport.topleftImageY+configuration.Global.TileSize),
								).(*ebiten.Image), op)
							}

							if configuration.Global.ActiveParticlesPortal {
								particle.Draw(screen, teleport.particles, posX, posY)
							}

						}
					}
				}
			}
		} else {
			// on vérifie qu'il existe
			if teleport.exist {
				// on créer les coordonnée par rapport au content
				var relatifX int = teleport.X - camXPos + configuration.Global.ScreenCenterTileX
				var relatifY int = teleport.Y - camYPos + configuration.Global.ScreenCenterTileY

				// on affiche si le téléporteur est dans l'écran
				var condition_affichage bool
				if configuration.Global.CameraMode == 2 {
					condition_affichage = relatifX >= -1 && relatifY >= -1 && relatifX <= configuration.Global.NumTileX && relatifY <= configuration.Global.NumTileY
				} else {
					condition_affichage = relatifX >= 0 && relatifY >= 0 && relatifX < configuration.Global.NumTileX && relatifY < configuration.Global.NumTileY

				}
				if condition_affichage {
					// on place l'image
					op := &ebiten.DrawImageOptions{}
					var posX float64
					var posY float64

					if FluideX {
						posX = float64((relatifX)*configuration.Global.TileSize - Xshift)
					} else {
						posX = float64((relatifX) * configuration.Global.TileSize)
					}

					if FluideY {
						posY = float64((relatifY)*configuration.Global.TileSize - Yshift)
					} else {
						posY = float64((relatifY) * configuration.Global.TileSize)
					}
					op.GeoM.Translate(posX, posY)

					// on découpe l'image pour que le portail disparaisse progressivement
					if (relatifX*configuration.Global.TileSize-Xshift) >= (configuration.Global.NumTileX-1)*configuration.Global.TileSize && (relatifY*configuration.Global.TileSize-Yshift) >= (configuration.Global.NumTileY-1)*configuration.Global.TileSize {
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize-((relatifX*configuration.Global.TileSize-Xshift)-(configuration.Global.NumTileX-1)*configuration.Global.TileSize), teleport.topleftImageY+configuration.Global.TileSize-((relatifY*configuration.Global.TileSize-Yshift)-(configuration.Global.NumTileY-1)*configuration.Global.TileSize)),
						).(*ebiten.Image), op)
					} else if (relatifX*configuration.Global.TileSize - Xshift) >= (configuration.Global.NumTileX-1)*configuration.Global.TileSize {
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize-((relatifX*configuration.Global.TileSize-Xshift)-(configuration.Global.NumTileX-1)*configuration.Global.TileSize), teleport.topleftImageY+configuration.Global.TileSize),
						).(*ebiten.Image), op)
					} else if (relatifY*configuration.Global.TileSize - Yshift) >= (configuration.Global.NumTileY-1)*configuration.Global.TileSize {
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize, teleport.topleftImageY+configuration.Global.TileSize-((relatifY*configuration.Global.TileSize-Yshift)-(configuration.Global.NumTileY-1)*configuration.Global.TileSize)),
						).(*ebiten.Image), op)
					} else {
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(teleport.topleftImageX, teleport.topleftImageY, teleport.topleftImageX+configuration.Global.TileSize, teleport.topleftImageY+configuration.Global.TileSize),
						).(*ebiten.Image), op)
					}
					if configuration.Global.ActiveParticlesPortal {
						particle.Draw(screen, teleport.particles, posX, posY)
					}

				}
			}
		}
		if p.timeMessage > 0 {
			p.timeMessage--
			ebitenutil.DebugPrintAt(screen, "Teleportation impossible", 0, (configuration.Global.NumTileY-1)*configuration.Global.TileSize)
		}
	}

}
