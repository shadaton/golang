package main

import "fmt"

func main() {
	x := 10
	fmt.Println("value is %d\n", x)
	fmt.Println("Address x variable %x\n", &x)
	var p *int
	p = &x // ขี้ไปยังAddress ที่ x อยู่
	fmt.Printf("Pointer p = %x\n", p)
	*p = 20
	fmt.Printf("value is %d\n", x)
}
