package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Println("[강제 형변환을 진행한다.]")
	fmt.Println("[1개의 음수를 고르세요]")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("다시 선택해!")
	} else {
		var b = int(a)
		var c = byte(a)
		var d = uint(a)

		fmt.Printf("음수의 int : %d \n", b)
		fmt.Printf("음수의 byte : %d\n", c)
		fmt.Printf("음수의 unit : %d\n", d)
	}
}
