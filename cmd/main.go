package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/game"

	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/assets"
)

func main() {

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "select configuration file")

	flag.Parse()

	configuration.Load(configFileName)

	assets.Load()

	g := &game.Game{}
	//ebiten.RunGameWithOptions(g, &ebiten.RunGameOptions{ScreenTransparent: false})
	g.Init()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("Quadtree")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
