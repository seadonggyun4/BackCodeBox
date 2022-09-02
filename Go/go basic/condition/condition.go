package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Scanf("%d %d", &num1, &num2)

	var result1 = num1 / num2
	var result2 = num1 % num2

	fmt.Printf("몫 : %d, 나머지 : %d", result1, result2)
}
