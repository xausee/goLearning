package main

import "fmt"

func insertSort(m []int) {
	for i := 1; i < len(m); i++ {
		tmp, j := m[i], i-1
		for ; j >= 0 && tmp < m[j]; j-- {
			m[j+1] = m[j]
		}
		m[j+1] = tmp
		fmt.Println(i, m)
	}
}

func main() {
	a := []int{123, 43, 5, 78, 89, 9, 0, 4}
	fmt.Println("original array: \n", a)
	insertSort(a)
	fmt.Println("after sorted: \n", a)
}
