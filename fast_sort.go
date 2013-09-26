package main

import (
	"fmt"
)

func quickSort(m []int, start int, end int) {
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
		quickSort(m, start, j-1)
		quickSort(m, j+1, end)
	}
}

func QuickSort(m []int) {
	quickSort(m, 0, len(m)-1)
}

func main() {
	a := []int{10, 3, 56, 67, 4, 33, 123, 0, 7, 50}
	fmt.Println("oringinal array: \n", a)
	QuickSort(a)
	fmt.Println("after sorted:  \n", a)
}
