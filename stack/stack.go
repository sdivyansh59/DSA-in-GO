package stack

// Stack represent stack that hold a slice
type Stack struct {
	item [] int
}

// Push  will push elelemt in stack
func ( s * Stack) Push ( i int) {
	s.item = append(s.item, i)
}

// Pop will remove element from the end and 
// return the poped element
func (s *Stack) Pop() int {
	l := len(s.item)-1
	// if stack is empty return -1
	if l <0 {
		return -1
	}
	toPop := s.item[l]
	s.item = s.item[:l]
	return toPop
}

// IsEmpty
func (s * Stack) IsEmpty() bool {
	if len(s.item) == 0 {
		return true
	}else {
		return false
	}
	// go don't have ternary opearator ?:
}
