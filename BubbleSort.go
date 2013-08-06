package main

import "fmt"

/* from start to end
   flag = true ascending order
   flag = false descending order */
func BubbleSortOne(m []int, flag bool) {  
  b := false   
  o: for i := 0; i < len(m) - 1; i++ {      
    if b {
      break o
    }
    b = true 
    for j := i + 1; j < len(m); j++ {           
      if flag {
        if m[i] > m[j] {          
          b, m[i], m[j] = false, m[j], m[i] 
        } 
      } else {
          if m[i] < m[j] {            
            b, m[i], m[j] =false, m[j], m[i] 
          }
      }      
    }
    fmt.Println(i,m) 
  }   
}

/* from end to start
   flag = true ascending order
   flag = false descending order */
func BubbleSortTwo(m []int, flag bool) { 
  b := false        
  o: for i := 0; i < len(m); i++ {     
    if b {
      break o
    }
    b = true  
    for j := 0; j < len(m)-i-1; j++ {
      if flag {
        if m[j] > m[j+1] {
          b, m[j+1], m[j] = false, m[j], m[j+1] 
        }
       } else {
           if m[j] < m[j+1] {
             b, m[j+1], m[j] =false, m[j], m[j+1] 
           }
       }           
    } 
    fmt.Println(i, m)      
  }  
}

func main() {
  a := []int{45, 1, 1000, 3, 4, 2, 5, 777, 43, 6, 23} 
  fmt.Println()
  fmt.Println("Oringinal array:\n",a)
  fmt.Println()
  BubbleSortOne(a, true)
  fmt.Println()
  BubbleSortOne(a, false)
  fmt.Println()
  BubbleSortTwo(a, true)
  fmt.Println()
  BubbleSortTwo(a, false)  
}
