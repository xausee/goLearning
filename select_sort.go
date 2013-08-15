package main

import "fmt"

func selectSort(m []int) {
	for i := 0; i < len(m); i++ {
		k := i
		for j := i + 1; j < len(m); j++ {
			if m[j] < m[k] {
				k = j
			}
			if k != i {
				m[i], m[k] = m[k], m[i]
			}
		}
		fmt.Println(i, m)
	}
}

func main() {
	a := []int{10, 3, 56, 67, 4, 33, 123, 0, 7, 50}
	fmt.Println("oringinal array: \n", a)
	selectSort(a)
	fmt.Println("after sorted:  \n", a)
}
