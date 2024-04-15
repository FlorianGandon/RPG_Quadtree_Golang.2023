package floor

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/generation"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/quadtree"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"

	"bufio"
	"log"
	"os"
	"strconv"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.Content = make([][]tiles.Tiles, configuration.Global.NumTileY+2)
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]tiles.Tiles, configuration.Global.NumTileX+2)
	}

	switch configuration.Global.FloorKind {
	case gridFloor:
		// on ne touche pâs à full content
	case fromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
		f.Height = len(f.fullContent)
		f.Width = len(f.fullContent[0])
	case quadTreeFloor:
		if !configuration.Global.RandomGeneration {
			fullContent := readFloorFromFile(configuration.Global.FloorFile)
			f.Height = len(fullContent)
			f.Width = len(fullContent[0])
			f.QuadtreeContent = quadtree.MakeFromArray(fullContent)
		} else {
			f.QuadtreeContent, f.ListChunkgenerate = generation.Init()
		}
	default:
		panic("unhandled default case")
	}
}

// Lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]tiles.Tiles) {
	// on définit les variables :

	// ligne_string représente chage ligne sous forme d'une chaine de caractère
	var ligne_string string
	// ligne_tab_int représente chage ligne sous forme d'un tableau de tuiles
	var ligne_tab_int []tiles.Tiles
	var err error
	var file *os.File

	// Ouverture du fichier
	file, err = os.Open(fileName)
	//on vérifie que le fichier existe
	if err != nil {
		log.Fatal(err)
	}

	var scanner *bufio.Scanner = bufio.NewScanner(file)

	// Lecture des lignes du fichier
	for scanner.Scan() {
		ligne_string = scanner.Text()
		ligne_tab_int = make([]tiles.Tiles, len(ligne_string))
		// pour les chiffres sous le format string dans ligne_string on les transforme en int et on les ajoutes au tableau
		for i := 0; i < len(ligne_string); i++ {
			ligne_tab_int[i].Types, err = strconv.Atoi(string(ligne_string[i]))
			// si un caractère n'est pas un chiffre
			if err != nil {
				log.Fatal(err)
			}
		}
		// on ajoute la ligne à floorContent
		floorContent = append(floorContent, [][]tiles.Tiles{ligne_tab_int}...)
	}
	// Vérification que tout s'est bien passé
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Fermeture du fichier
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	return floorContent
}
