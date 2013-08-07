package main

import "fmt"

type Stack struct {
  i int 
  data [10]int
}

func (s *Stack) Size() (int){
	return s.i	
}

func (s *Stack) Push(v int) {
  if s.i > 9 { 
  	fmt.Printf("stack is full, fail to push value: %v\n", v)
  	return 
  }
  s.data[s.i] = v
  s.i++
}

func (s *Stack) Pop() (ret int) {  
  s.i--
  if s.i < 0 {
    fmt.Printf("stack is empty! ") 
  	s.i = 0; 
  	return 
  }
  ret, s.data[s.i] = s.data[s.i], 0  
  return
}

func TestPushPop() {
  s := new(Stack)
  for i := 0; i < 15; i++ {
  	s.Push(i)  	
  }
  fmt.Printf("stack size is: %v\nstack content is: %v\n", s.Size(), s)
  for i := 0; i < 15; i++ {
  	fmt.Printf("pop up value: %v\n", s.Pop())  	
  }   
  fmt.Printf("stack size is: %v\nstack content is: %v\n", s.Size(), s)
}

func main(){	
	TestPushPop()
}