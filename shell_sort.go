package main

import "fmt"

func shellSort(m []int) {
	for d := len(m) / 2; d >= 1; d = d / 2 {
		for i := d; i < len(m); i++ {
			tmp, j := m[i], i-d
			for ; j >= 0 && tmp < m[j]; j -= d {
				m[j+d] = m[j]
			}
			m[j+d] = tmp
			fmt.Println(d, m)
		}
	}
}

func main() {
	a := []int{123, 43, 5, 78, 89, 9, 0, 4}
	fmt.Println("original array: \n", a)
	shellSort(a)
	fmt.Println("after sorted: \n", a)
}
