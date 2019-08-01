package main

import "fmt"

func main() {
	dosomething("shadaton")
	dosomething("TUTORIAL")
	addition(8, 2)
	result := addition2(5, 5)
	fmt.Println(result)
}

//ไม่มีการคืนค่าและรับ
func dosome() {
	fmt.Println("OK")
}

//รับค่า
func dosomething(str string) {
	fmt.Println(str)
}

//รับค่า
func addition(a int, b int) {
	fmt.Println(a + b)
}

//คืนค่า return คืนค่า
func addition2(a int, b int) int {
	output := (a + b)
	return output
}
