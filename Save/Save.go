package Save

import (
	"encoding/json"
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/Coords"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// SaveMap sauvegarde la carte dans un fichier .map
func SaveMap(maMap map[Coords.Coords][][]tiles.Tiles, nomFichier string, x, y int) error {
	// Encoder la carte en JSON
	var Map = make(map[string][][]tiles.Tiles, 10)
	for i, y := range maMap {
		Map[CoordsToString(i)] = y
	}
	jsonData, err := json.Marshal(Map)
	if err != nil {
		return err
	}

	// Écrire les données JSON dans le fichier
	err = os.WriteFile(nomFichier, jsonData, 0644)
	if err != nil {
		return err
	}
	// Encoder la config en JSON
	configuration.Global.PlayerX = x
	configuration.Global.PlayerY = y
	config, err := json.Marshal(configuration.Global)
	if err != nil {
		return err
	}
	err = os.WriteFile(strings.TrimSuffix(nomFichier, filepath.Ext(nomFichier))+".json", config, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Carte sauvegardée avec succès.")
	return nil
}

func CoordsToString(coord Coords.Coords) (f string) {
	return strconv.Itoa(coord.X) + "," + strconv.Itoa(coord.Y)
}

func StringtoCoords(s string) (c Coords.Coords) {
	part := strings.Split(s, ",")
	fmt.Println(part)
	var x, _ = strconv.Atoi(part[0])
	var y, _ = strconv.Atoi(part[1])
	return Coords.Coords{X: x, Y: y}
}

// Charge la map depuis un fichier .json
func LoadMap(nomFichier string) (maps map[Coords.Coords][][]tiles.Tiles, err error, x int, y int) {
	// Lire les données depuis le fichier
	jsonData, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err, 0, 0
	}

	// Décoder les données JSON dans une nouvelle carte
	var maMap map[string][][]tiles.Tiles
	err = json.Unmarshal(jsonData, &maMap)
	if err != nil {
		return nil, err, 0, 0
	}
	var Map = make(map[Coords.Coords][][]tiles.Tiles, 10)
	for i, y := range maMap {
		Map[StringtoCoords(i)] = y
	}
	configuration.Load(strings.TrimSuffix(nomFichier, filepath.Ext(nomFichier)) + ".json")
	x = configuration.Global.PlayerX
	y = configuration.Global.PlayerY
	fmt.Println("Carte chargée avec succès.", x, y)
	return Map, nil, x, y
}
