package main

import "fmt"

func fibonacci(value int) ([]int){
  m := make([]int, value)
  m[0], m[1] = 1, 1
  for n := 2; n < value; n++ {
    m[n] = m[n -1] + m[n - 2]
  }
  return m  
}

func main() {
  for _, term := range fibonacci(20) {
    fmt.Println(term)
  }
}
