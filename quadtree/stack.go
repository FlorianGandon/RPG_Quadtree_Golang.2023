package quadtree

// Creation d'une structure de pile
// top l'element au sommet
// height le nombre d'element
type stack struct {
	top    *element
	height int
}

// un element pointe vers le suivant dans la pille
// et il pointe ver un noeud
type element struct {
	n    *node
	next *element
}

// ajouter un élément dans une pile
func (s *stack) push(e *element) {
	if s.top == nil {
		s.top = e
		s.height = 1
	} else {
		e.next = s.top
		s.top = e
		s.height++
	}
}

// retirer un élément dans une pile
func (s *stack) pop() (e *element) {
	if s.top != nil {
		e = s.top
		s.top = e.next
		s.height--
		return e
	}
	return nil
}

// vérifier si la pile est vide
func (s *stack) is_empty() bool {
	return s.height == 0
}
