package container

type Stack struct {
	buff []interface{}
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(v interface{}) {
	if s.size == len(s.buff) {
		s.buff = append(s.buff, v)
	} else {
		s.buff[s.size] = v
	}
	s.size++
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.size == 0 {
		return nil, false
	}
	s.size--
	return s.buff[s.size], true
}
