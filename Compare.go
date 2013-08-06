package main

import "fmt"

func compare(a, b []byte) (int){
	for i := 0; i < len(a) && i < len(b); i++{
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}

		switch {
		case len(a) > len(b):
			return 1
		case len(a) < len(b):
			return -1
		}
	}
	return 0
}

func main(){
	a := []byte{'h','e','l','l','o',','}
	b := []byte{'w','o','r','l','d','!'}
	c := []byte{'w','o','r','l','d','!'}
	d := []byte{'w','o','r','l','d','!', 'h', 'e'}
	fmt.Println(compare(a, b), compare(b, c), compare(b, d), compare(b, a))
}