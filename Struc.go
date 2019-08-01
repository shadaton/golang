package main

import (
	"fmt"
)

type Books struct {
	title    string
	author   string
	subtitle string
	price    float32
}

func main() {
	Golang := Books{title: "Go Programing", author: "KongRksiam", price: 200}
	fmt.Println(Golang)
}
