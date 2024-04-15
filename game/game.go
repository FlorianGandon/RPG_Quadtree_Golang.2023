package game

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/camera"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/character"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/floor"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/portal"
)

// Game est le type permettant de représenter les données du jeu.
// Aucun champs n'est exporté pour le moment.
//
// Les champs non exportés sont :
//   - camera : la représentation de la caméra
//   - floor : la représentation du terrain
//   - character : la représentation du personnage
type Game struct {
	camera    camera.Camera
	floor     floor.Floor
	character character.Character
	portal    portal.Portal
}
