package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("a의 값 입력:")
	fmt.Scanln(&a)
	fmt.Print("b의 값 입력:")
	fmt.Scanln(&b)

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d X %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d 의 나머지 : %d\n", a, b, a%b)

}
