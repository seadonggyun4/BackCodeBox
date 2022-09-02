package main

import "fmt"

func main() {
	var A [10]int
	//len 배열 크기 표출
	for i := 0; i < len(A); i++ {
		A[i] = i + i
	}

	fmt.Println(A)
}
