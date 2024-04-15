package game

import (
	"errors"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/Save"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/generation"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sqweek/dialog"
	"log"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {
	// la touche D active le debug mode
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}
	if configuration.Global.ActiveScrolling {
		_, dy := ebiten.Wheel()
		if dy > 0 || (!(configuration.Global.NumTileX <= 3) && dy < 0) {
			configuration.Global.NumTileX += int(dy) * 2
			configuration.Global.NumTileY += int(dy) * 2
			configuration.Global.ScreenWidth = configuration.Global.NumTileX * configuration.Global.TileSize
			configuration.Global.ScreenHeight = configuration.Global.NumTileY * configuration.Global.TileSize
			configuration.Global.ScreenCenterTileX += int(dy)
			configuration.Global.ScreenCenterTileY += int(dy)
			g.floor.Content = make([][]tiles.Tiles, configuration.Global.NumTileY+2)
			for y := 0; y < len(g.floor.Content); y++ {
				g.floor.Content[y] = make([]tiles.Tiles, configuration.Global.NumTileX+2)
			}
		}
	}
	// on actualise le joueur en fonction du debug mode
	if configuration.Global.DebugMode {
		g.character.Update([4]bool{false, false, false, false}, g.floor.Width, g.floor.Height)
	} else {
		g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y), g.floor.Width, g.floor.Height)
	}
	if configuration.Global.ActivePortal {
		// la touche T permet de se téléporter
		if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			possible, newX, newY := g.portal.CanUseTeleporter(g.character.X, g.character.Y)
			if possible {
				g.character.Teleport(newX, newY)
				g.camera.Teleport(newX, newY, g.floor.Width, g.floor.Height)
			}
		}
	}

	if inpututil.KeyPressDuration(ebiten.KeyControl) > 0 && inpututil.KeyPressDuration(ebiten.KeyUp) > 0 && (configuration.Global.RandomGeneration || configuration.Global.FloorKind == 0) && configuration.Global.ActiveBigStep {
		g.character.Teleport(g.character.X, g.character.Y-configuration.Global.BigStep)
		g.camera.Teleport(g.character.X, g.character.Y-configuration.Global.BigStep, g.floor.Width, g.floor.Height)
	}
	if inpututil.KeyPressDuration(ebiten.KeyControl) > 0 && inpututil.KeyPressDuration(ebiten.KeyDown) > 0 && (configuration.Global.RandomGeneration || configuration.Global.FloorKind == 0) && configuration.Global.ActiveBigStep {
		g.character.Teleport(g.character.X, g.character.Y+configuration.Global.BigStep)
		g.camera.Teleport(g.character.X, g.character.Y+configuration.Global.BigStep, g.floor.Width, g.floor.Height)
	}
	if inpututil.KeyPressDuration(ebiten.KeyControl) > 0 && inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 && (configuration.Global.RandomGeneration || configuration.Global.FloorKind == 0) && configuration.Global.ActiveBigStep {
		g.character.Teleport(g.character.X-configuration.Global.BigStep, g.character.Y)
		g.camera.Teleport(g.character.X-configuration.Global.BigStep, g.character.Y, g.floor.Width, g.floor.Height)
	}
	if inpututil.KeyPressDuration(ebiten.KeyControl) > 0 && inpututil.KeyPressDuration(ebiten.KeyRight) > 0 && (configuration.Global.RandomGeneration || configuration.Global.FloorKind == 0) && configuration.Global.ActiveBigStep {
		g.character.Teleport(g.character.X+configuration.Global.BigStep, g.character.Y)
		g.camera.Teleport(g.character.X+configuration.Global.BigStep, g.character.Y, g.floor.Width, g.floor.Height)
	}
	if configuration.Global.RandomGenerationSave {
		if configuration.Global.RandomGeneration && inpututil.KeyPressDuration(ebiten.KeyControl) != 0 && inpututil.IsKeyJustPressed(ebiten.KeyS) {
			filename, err := dialog.File().Title("Sauvegarder une carte").Filter("Carte", "map").SetStartDir("./maps").SetStartFile("MaCarte.map").Save()
			if err != nil || errors.Is(err, dialog.ErrCancelled) {
				log.Println(err)
			} else {
				err := Save.SaveMap(g.floor.ListChunkgenerate, filename, g.character.X, g.character.Y)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if configuration.Global.RandomGeneration && inpututil.KeyPressDuration(ebiten.KeyControl) != 0 && inpututil.IsKeyJustPressed(ebiten.KeyO) {
			filename, err := dialog.File().Title("Ouvrir une carte").Filter("carte", "map").SetStartDir(".").Load()
			if err != nil || errors.Is(err, dialog.ErrCancelled) {
				log.Println(err)
			} else {
				g.floor.ListChunkgenerate, err, g.character.X, g.character.Y = Save.LoadMap(filename)
				if err != nil {
					log.Fatal(err)
				}
				g.camera.X, g.camera.Y = 0, 0
			}

		}
	}

	g.camera.Update(g.character.X, g.character.Y, g.character.Orientation, g.floor.Width, g.floor.Height)
	configuration.Global.CameraX = g.camera.X
	configuration.Global.CameraY = g.camera.Y
	g.floor.Update(g.camera.X, g.camera.Y)
	if configuration.Global.RandomGeneration {
		g.floor.QuadtreeContent, g.floor.GX, g.floor.GY = generation.Update(g.floor.ListChunkgenerate)
	}
	if configuration.Global.ActivePortal {
		// la touche I permet de poser / supprimer les portails
		if inpututil.IsKeyJustPressed(ebiten.KeyI) {
			// on vérifie si le joueur est dans la carte
			if (g.character.X >= 0 && g.character.X < g.floor.Width && g.character.Y >= 0 && g.character.Y < g.floor.Height) || configuration.Global.RandomGeneration || configuration.Global.DebugMode {
				g.portal.Interact(g.character.X, g.character.Y)
			}
		}
	}

	g.portal.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		configuration.Global.FullScreen = !configuration.Global.FullScreen
		ebiten.SetFullscreen(configuration.Global.FullScreen)
	}
	return nil
}
