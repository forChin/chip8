package main

type stack []word // word ?

func (s *stack) pop() word {
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return elem
}

// return bool ?
func (s *stack) push(elem word) {
	*s = append(*s, elem)
}
