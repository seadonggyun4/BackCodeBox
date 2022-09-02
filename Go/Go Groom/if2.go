package main

import "fmt"

func main() {
	fmt.Println("[1~100 사이의 5와 8의 배수]")

	for i := 1; i < 101; i++ {
		if (i%5 == 0) || (i%8 == 0) || ((i%5 == 0) && (i%8 == 0)) {
			fmt.Printf("5 or 8의 배수: %d \n", i)
		}

	}
}
