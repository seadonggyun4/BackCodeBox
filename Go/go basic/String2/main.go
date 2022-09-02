package main

import "fmt"

func main() {
	s := "Hello 월드"

	s2 := []rune(s)

	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ",")

	}
}
