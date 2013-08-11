package main

import "fmt"

func insert_sort(m []int) {
	for i := 1; i < len(m); i++ {
		tmp := m[i]
		j := i - 1
		for tmp < m[j] && j > 0 {
			m[j+1] = m[j]
			fmt.Println(j)
			j--
		}
		m[j+1] = tmp
		fmt.Println(m)
	}
}

func main() {
	a := []int{123, 43, 5, 78, 89, 9, 0, 4}
	fmt.Println("original array: \n", a)
	insert_sort(a)
	fmt.Println("after sorted: \n", a)
}
