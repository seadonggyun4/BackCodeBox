package main

import (
	"fmt"
)

func main() {
	number1, number2 := 1, 2

	number2++    // 증감연산자는 뒤에만 가능
	number1 += 3 //3+1 = 4 할당연산자
	number1 %= 2 // 4/2=0 0을 대입
	fmt.Println(number2)
	fmt.Println(number1)

	//=====================할당 연산자 ===========================
	num1, num2 := 1, 2
	num2 <<= 2 // 쉬프트 연산자: 화살표 방향으로 정해진 수만큼 이동 -> 2진수 화해서
	fmt.Printf("%08b <<= 2 의 값:  %08b, %d \n\n", 2, num2, num2)

	num1 &= 1 //And 연산자 : 10진수를 2진수로 바꿔 연산하고 같은 값만 같게 표출
	fmt.Printf("%08b &= %08b 의 값 : %08b, %d, \n", 1, 1, num1, num1)

	// ===============비트 연산자 ===============
	num4 := 1
	fmt.Printf("%08b & %08b 의 값 : %08b  %d \n\n", 1, 1, num4&1, num4&1)

	// ================== 논리연산자 ===========
	// boo1 타입으로 만 연산 가능
	num3 := true
	fmt.Println("num3 && true 의 값 :", num3 && true)

}
