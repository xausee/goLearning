package main

import (
	"fmt"
)

func fastSort(m []int, start int, end int) {
	if start < end {
		i, j, x := start, end, m[start]
		for i < j {
			for m[j] >= x && j > i {
				j--
			}
			if i < j {
				m[i], m[j] = m[j], m[i]
				i++
			}
			for m[i] <= x && j > i {
				i++
			}
			if i < j {
				m[i], m[j] = m[j], m[i]
				j--
			}
		}
		fmt.Println(m)
		fastSort(m, start, j-1)
		fastSort(m, j+1, end)
	}
}

func main() {
	a := []int{10, 3, 56, 67, 4, 33, 123, 0, 7, 50}
	fmt.Println("oringinal array: \n", a)
	fastSort(a, 0, len(a)-1)
	fmt.Println("after sorted:  \n", a)
}
