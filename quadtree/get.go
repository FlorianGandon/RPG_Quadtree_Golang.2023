package quadtree

import "github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]tiles.Tiles) {
	var s stack = stack{}
	var current_element *element
	var width int = len(contentHolder[0])
	var height int = len(contentHolder)

	// décalage pour charger une case de plus tout autour de content
	topLeftX--
	topLeftY--

	s.push(&element{n: q.Root})

	// boucle pour parcourir l'arbre avec une pile tant qu'elle n'est pas vide
	for !s.is_empty() {
		// on récupère l'élément en haut de la pile
		current_element = s.pop()
		if current_element.n.content == -1 {
			// on regarde pour chaque noeud si il est dans la zone du content, si oui on l'ajoute à la pile
			if current_element.n.bottomRightNode != nil && (current_element.n.bottomRightNode.topLeftX <= (topLeftX + width - 1)) && ((current_element.n.bottomRightNode.topLeftX + current_element.n.bottomRightNode.width - 1) >= topLeftX) && (current_element.n.bottomRightNode.topLeftY <= (topLeftY + height - 1)) && ((current_element.n.bottomRightNode.topLeftY + current_element.n.bottomRightNode.height - 1) >= topLeftY) {
				s.push(&element{n: current_element.n.bottomRightNode})
			}
			if current_element.n.bottomLeftNode != nil && (current_element.n.bottomLeftNode.topLeftX <= (topLeftX + width - 1)) && ((current_element.n.bottomLeftNode.topLeftX + current_element.n.bottomLeftNode.width - 1) >= topLeftX) && (current_element.n.bottomLeftNode.topLeftY <= (topLeftY + height - 1)) && ((current_element.n.bottomLeftNode.topLeftY + current_element.n.bottomLeftNode.height - 1) >= topLeftY) {
				s.push(&element{n: current_element.n.bottomLeftNode})
			}
			if current_element.n.topRightNode != nil && (current_element.n.topRightNode.topLeftX <= (topLeftX + width - 1)) && ((current_element.n.topRightNode.topLeftX + current_element.n.topRightNode.width - 1) >= topLeftX) && (current_element.n.topRightNode.topLeftY <= (topLeftY + height - 1)) && ((current_element.n.topRightNode.topLeftY + current_element.n.topRightNode.height - 1) >= topLeftY) {
				s.push(&element{n: current_element.n.topRightNode})
			}
			if current_element.n.topLeftNode != nil && (current_element.n.topLeftNode.topLeftX <= (topLeftX + width - 1)) && ((current_element.n.topLeftNode.topLeftX + current_element.n.topLeftNode.width - 1) >= topLeftX) && (current_element.n.topLeftNode.topLeftY <= (topLeftY + height - 1)) && ((current_element.n.topLeftNode.topLeftY + current_element.n.topLeftNode.height - 1) >= topLeftY) {
				s.push(&element{n: current_element.n.topLeftNode})
			}
		} else {
			var debutX int
			var debutY int
			var finX int
			var finY int

			// on affiche le noeud

			if current_element.n.topLeftX >= topLeftX {
				debutX = current_element.n.topLeftX - topLeftX
			} else {
				debutX = 0
			}
			if current_element.n.topLeftY >= topLeftY {
				debutY = current_element.n.topLeftY - topLeftY
			} else {
				debutY = 0
			}
			if (current_element.n.topLeftX + current_element.n.width - 1) >= (topLeftX + width - 1) {
				finX = width - 1
			} else {
				finX = (current_element.n.topLeftX + current_element.n.width - 1) - topLeftX
			}
			if (current_element.n.topLeftY + current_element.n.height - 1) >= (topLeftY + height - 1) {
				finY = height - 1
			} else {
				finY = (current_element.n.topLeftY + current_element.n.height - 1) - topLeftY
			}
			for y := debutY; y <= finY; y++ {
				for x := debutX; x <= finX; x++ {
					contentHolder[y][x].Types = current_element.n.content
				}
			}
		}
	}
}
