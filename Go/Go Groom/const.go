package main

import "fmt"

const (
	a = 1
	b = 2
)

func main() {

	const (
		c = 1
		d = "Hi"
	)

	const e float32 = 3.32

	fmt.Print(a, b, c, d, e)
}

/*
cost 는 변경할수 없는숫자인 상수!
상수는 ()를 통해 한번에 여러가지와 여러 타입 선언가능
*/
