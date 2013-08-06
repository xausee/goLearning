package main

import "fmt"

func bubblesortone(m []int, flag bool) {     
  for i := 0; i < len(m) - 1; i++ {    
    for j := i + 1 ; j < len(m); j++ {
      if flag {
        if m[i] > m[j] {
          m[i], m[j] = m[j], m[i] 
        } 
      } else {
          if m[i] < m[j] {
          m[i], m[j] = m[j], m[i] 
          }
      }       
    }  
  }  
}

func bubblesorttwo(m []int, flag bool) {        
  for i := len(m) - 1; i > 0; i-- {    
    for j := 0; j < i; j++ {
      if flag {
        if m[j] > m[i] {
         m[i], m[j] = m[j], m[i]
        }
       } else {
           if m[j] < m[i] {
             m[i], m[j] = m[j], m[i]
           }
       }
    }   
  }  
}

func main() {
  a := []int{1, 1000, 3, 4, 2, 5, 777, 43, 6, 23}
  b := []int{1, 1000, 3, 4, 2, 5, 777, 43, 6, 23}
  fmt.Println("unsorted array: ", a, b)
  bubblesortone(a, false)
  bubblesorttwo(b, true)
  fmt.Println("sorted array: ", a, b) 
}
