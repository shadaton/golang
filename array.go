package main

import "fmt"

func main() {
	x := [5]float64{10, 20, 30, 40, 50}
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
