package main

type stack []uint16 // word ?

func (s *stack) pop() uint16 {
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return elem
}

// return bool ?
func (s *stack) push(elem uint16) {
	*s = append(*s, elem)
}
