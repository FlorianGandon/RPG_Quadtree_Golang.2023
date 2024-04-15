package game

import "github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.character.Init()
	g.floor.Init()
	g.camera.Init(g.floor.Width, g.floor.Height)
	configuration.Global.CameraX = g.camera.X
	configuration.Global.CameraY = g.camera.Y
	g.portal.Init()
}
