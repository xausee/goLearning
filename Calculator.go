package main

import (
  "fmt"  
  "bufio" 
  "os"
  "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)

type Stack struct {
  i int 
  data [10]int
}

func (s *Stack) push(v int) {
  if s.i +1 > 9 { return }
  s.data[s.i] = v
  s.i++
}

func (s *Stack) pop() (ret int) {  
  s.i--
  if s.i < 0 { s.i = 0; return }
  ret = s.data[s.i]
  return
}

func main() {
        for {
               s, err := reader.ReadString('\n')
               var token string
               if err != nil { return }
               for _, c := range s {
                       switch {
                       case c >= '0' && c <= '9':
                           token = token + string(c)
                       case c == ' ':
                           r, _ := strconv.Atoi(token)
                           st.push(r)                           
                           token =""
                       case c == '+':
                           fmt.Printf("%d\n", st.pop() + st.pop()) 
                       case c == '*':
                           fmt.Printf("%d\n", st.pop() * st.pop()) 
                       case c == '-':
                           fmt.Printf("%d\n", st.pop() - st.pop())
                       case c == '/':
                           p := st.pop()
                           q := st.pop()
                           if q == 0 { return }
                           fmt.Printf("%d\n", p / q) 
                       case c == '%':
                           p := st.pop()
                           q := st.pop()
                           if q == 0 { return }
                           fmt.Printf("%d\n", p % q)   
                       case c == 'q':
                           return 
                       default:
                           // err   
                       }             
               }
        }
}
