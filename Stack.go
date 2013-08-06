package stack

type Stack struct {
  i int 
  data [10]int
}

func (s *Stack) Push(v int) {
  if s.i +1 > 9 { return }
  s.data[s.i] = v
  s.i++
}

func (s *Stack) Pop() (ret int) {  
  s.i--
  if s.i < 0 { s.i = 0; return }
  ret = s.data[s.i]
  return
}
