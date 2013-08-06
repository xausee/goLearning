package main

import (
"fmt"
)

func  reverse(str string) (string) {
	var a = []rune(str)
	for i, j := 0, len(a)-1; i<j; i, j = i+1, j-1{
		a[i], a[j] = a[j], a[i]
	}
	var res = string(a)
	return res	
}

func main(){
	s0 := "hello world!"
	fmt.Println(s0)	
	var s1 = reverse(s0)
	fmt.Println(s1)
}
