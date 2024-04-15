package assets

import (
	"bytes"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
	"image"
	"log"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed terrain2.png
var floorBytes []byte

// FloorImage contient une version compatible avec Ebitengine de l'image
// qui contient les différents éléments qui peuvent s'afficher au sol
// (herbe, sable, etc).
// Dans la version du projet qui vous est fournie, ces éléments sont des
// carrés de 16 pixels de côté. Vous pourrez changer cela si vous le voulez.
var FloorImage *ebiten.Image

//go:embed character.png
var characterBytes []byte

//go:embed character2.png
var character2Bytes []byte

// CharacterImage contient une version compatible avec Ebitengine de
// l'image qui contient les différentes étapes de l'animation du
// personnage.
// Dans la version du projet qui vous est fournie, ce personnage tient
// dans un carré de 16 pixels de côté. Vous pourrez changer cela si vous
// le voulez.
var CharacterImage *ebiten.Image

// Load est la fonction en charge de transformer, à l'exécution du programme,
// les images du jeu en structures de données compatibles avec Ebitengine.
// Ces structures de données sont stockées dans les variables définies ci-dessus.
func Load() {
	decoded, _, err := image.Decode(bytes.NewReader(floorBytes))
	if err != nil {
		log.Fatal(err)
	}
	FloorImage = ebiten.NewImageFromImage(decoded)
	if configuration.Global.BetterCharacter {
		decoded, _, err = image.Decode(bytes.NewReader(character2Bytes))
	} else {
		decoded, _, err = image.Decode(bytes.NewReader(characterBytes))
	}
	if err != nil {
		log.Fatal(err)
	}
	CharacterImage = ebiten.NewImageFromImage(decoded)
}
