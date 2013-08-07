package main

import "fmt"

func Average(a []float64) (avr float64){
	sum := 0.0
	switch len(a){
	case 0:
		avr = 0
	default:
		for _, v := range a{
			sum += v
		}
		avr = sum / float64(len(a))
	}	
	return 
}

func main(){
	a := []float64{1, 2, 4.5}
	fmt.Println(Average(a))
}