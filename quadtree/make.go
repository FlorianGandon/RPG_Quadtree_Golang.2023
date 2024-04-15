package quadtree

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
)

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]tiles.Tiles) (q Quadtree) {
	q.Width = len(floorContent[0])
	q.Height = len(floorContent)

	q.Root = Build_quadtree(floorContent, &q)

	return
}

func Build_quadtree(floorContent [][]tiles.Tiles, quadtree *Quadtree) (root *node) {
	// si il n'y a pas de terrain, il n'y a pas de racine
	if floorContent == nil {
		return nil
	}
	root = &node{topLeftX: 0, topLeftY: 0, width: quadtree.Width, height: quadtree.Height, content: -1}
	// on vérifie si la racine est une feuille
	if estCarreIdentique(element{n: root}, floorContent) {
		root.content = floorContent[0][0].Types
		return root
	}
	var s stack //création d'une pile
	var current_element *element

	var middleX int = quadtree.Width - (quadtree.Width / 2)
	var middleY int = quadtree.Height - (quadtree.Height / 2)

	// si une de ses noeuds peut être créer on les ajoute à la pile

	// noeud bas droite
	if quadtree.Width-middleX > 0 && quadtree.Height-middleY > 0 {
		root.bottomRightNode = &node{topLeftX: middleX, topLeftY: middleY, width: quadtree.Width - middleX, height: quadtree.Height - middleY, content: -1}
		s.push(&element{n: root.bottomRightNode})
	}

	// noeud bas gauche
	if quadtree.Height-middleY > 0 {
		root.bottomLeftNode = &node{topLeftX: root.topLeftX, topLeftY: middleY, width: middleX, height: quadtree.Height - middleY, content: -1}
		s.push(&element{n: root.bottomLeftNode})
	}

	// noeud haut droite
	if quadtree.Width-middleX > 0 {
		root.topRightNode = &node{topLeftX: middleX, topLeftY: root.topLeftY, width: quadtree.Width - middleX, height: middleY, content: -1}
		s.push(&element{n: root.topRightNode})
	}
	// noeud haut gauche
	root.topLeftNode = &node{topLeftX: root.topLeftX, topLeftY: root.topLeftY, width: middleX, height: middleY, content: -1}
	s.push(&element{n: root.topLeftNode})
	// boucle pour parcourir l'arbre avec une pile tant qu'elle n'est pas vide
	for !s.is_empty() {
		// on récupère l'élément en haut de la pile
		current_element = s.pop()

		if estCarreIdentique(*current_element, floorContent) {
			current_element.n.content = floorContent[current_element.n.topLeftY][current_element.n.topLeftX].Types
		} else {

			middleX = current_element.n.width - (current_element.n.width / 2)
			middleY = current_element.n.height - (current_element.n.height / 2)

			// noeud bas droite
			if current_element.n.width-middleX > 0 && current_element.n.height-middleY > 0 {
				current_element.n.bottomRightNode = &node{topLeftX: current_element.n.topLeftX + middleX, topLeftY: current_element.n.topLeftY + middleY, width: current_element.n.width - middleX, height: current_element.n.height - middleY, content: -1}
				s.push(&element{n: current_element.n.bottomRightNode})
			}

			// noeud bas gauche
			if current_element.n.height-middleY > 0 {
				current_element.n.bottomLeftNode = &node{topLeftX: current_element.n.topLeftX, topLeftY: current_element.n.topLeftY + middleY, width: middleX, height: current_element.n.height - middleY, content: -1}
				s.push(&element{n: current_element.n.bottomLeftNode})
			}

			// noeud haut droite
			if current_element.n.width-middleX > 0 {
				current_element.n.topRightNode = &node{topLeftX: current_element.n.topLeftX + middleX, topLeftY: current_element.n.topLeftY, width: current_element.n.width - middleX, height: middleY, content: -1}
				s.push(&element{n: current_element.n.topRightNode})
			}
			// noeud haut gauche
			current_element.n.topLeftNode = &node{topLeftX: current_element.n.topLeftX, topLeftY: current_element.n.topLeftY, width: middleX, height: middleY, content: -1}
			s.push(&element{n: current_element.n.topLeftNode})

		}
	}
	return root
}

func estCarreIdentique(e element, maps [][]tiles.Tiles) bool {
	if e.n.width == 0 || e.n.height == 0 {
		// Si l'entrée est vide ou ne contient pas d'éléments, retourner false
		return false
	}

	var firstValue int = maps[e.n.topLeftY][e.n.topLeftX].Types // Prendre la première valeur comme référence

	for y := e.n.topLeftY; y < e.n.topLeftY+e.n.height; y++ {
		for x := e.n.topLeftX; x < e.n.topLeftX+e.n.width; x++ {
			if maps[y][x].Types != firstValue {
				// Si une valeur différente est trouvée, retourner false
				return false
			}
		}
	}
	// Si toutes les valeurs sont identiques, retourner true
	return true
}
