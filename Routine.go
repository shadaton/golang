package main

import "fmt"

func main() {
	go f(0)
	var input string
	fmt.Scan(&input)
}
func f(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, ":", i)
	}
}
