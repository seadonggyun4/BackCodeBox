package main

import "fmt"

func main() {
	s := "Hello World"

	temp01 := [11]byte{}
	temp02 := [11]rune{}
	run := []rune(s)
	temp03 := [11]string{}

	for i := 0; i < len(s); i++ {
		temp01[i] = s[(len(s)-1)-i]
	}

	for i := 0; i < len(s); i++ {
		temp02[i] = run[(len(s)-1)-i]
	}

	for i := 0; i < len(s); i++ {
		temp03[i] = string(s[(len(s)-1)-i])
	}

	fmt.Println(temp01)
	fmt.Println(temp02)
	fmt.Println(temp03)

}
