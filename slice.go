package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	silce2 := append(slice1, 9, 10)
	fmt.Println(silce2)
}
