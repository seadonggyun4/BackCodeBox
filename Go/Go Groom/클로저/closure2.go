package main

import "fmt"

func calc() func() int {
	result := 0
	return func() int {
		result += 1
		return result
	} // 반환 타입이 될 함수 내부 연산정의
}

func main() {

	calcA := calc()//매개변수가 선언안되어있는 함수라 그냥 변수 할당만 한다.
	fmt.Println(calcA())
	fmt.Println(calcA())
	fmt.Println(calcA())
	fmt.Println(calcA())

	calcB := calc()//새롭게 다시 선언
	fmt.Println(calcB())
	fmt.Println(calcB())
	fmt.Println(calcB())
	fmt.Println(calcB())

}