package main

import "fmt"

var c int = 1
var s = "Hello world!" // 반드시 형지정을 하지 않아도 된다.

func main() {
	a := 3 //이 선언방식은 초기화가 반드시 필요!, 또 함수 내부에서만 선언 가능
	b := 2

	fmt.Print(a, b, c)
	fmt.Println(a, ",", b, ",", c)
	fmt.Printf("%d,%d,%d \n", a, b, c)
	fmt.Println(s)
	fmt.Printf("%s", s)
} // 변수는 어떤 방식으로든 선언되면 반드시 사용되야한다!!!
