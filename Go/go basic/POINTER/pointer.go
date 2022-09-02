package main

import "fmt"

func main() {

	for i := 1; i < 101; i++ {

		if i%7 == 0 {
			fmt.Printf("%d", i)
		}

	}
}
