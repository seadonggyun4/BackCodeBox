/*
1. Go 언어에서 조건식(if/else문 )은 반드시 bool 형이다.

2. Go 언어에서는 c, c++과 다르게 bool 타입이 0,1 은 안되고 true, false만 가능

3. else 와 {} 를 같은 줄에 작성 하지 않으면 런타임 에러를 일으킨다!!

4. Go 언어 에서는 조건식 앞에 간단 한 문장을 넣어 실행이 가능하다.
=> 조건식 전에 정의된 변수는 해당 조건문 블록에서만 가능하다.
=> " if val := num*2; val == 2 " 라는 식이 있을때

*/

package main

import "fmt"

func main() {
	fmt.Println("[if 문을 활용한 계산기]")
	fmt.Println("")
	var num1, num2, num3 int //입력숫자1, 계산할숫자 2개 입력
	fmt.Println("계산 내용을 선택하세요")
	fmt.Println("1번: +")
	fmt.Println("2번: -")
	fmt.Println("3번: *")
	fmt.Println("4번: /")
	fmt.Print("=>")
	fmt.Scanln(&num1)
	fmt.Print("[계산할 두 숫자 입력] : ")
	fmt.Scanln(&num2, &num3)

	for {
		if num1 > 5 {
			fmt.Println("잘못된 선택입니다!")
		} else if num1 == 1 {
			fmt.Printf("%d + %d = %d \n", num2, num3, num2+num3)
			break
		} else if num1 == 2 {
			fmt.Printf("%d - %d = %d \n", num2, num3, num2-num3)
			break
		} else if num1 == 3 {
			fmt.Printf("%d X %d = %d \n", num2, num3, num2*num3)
			break
		} else if num1 == 4 {
			fmt.Printf("%d / %d = %d \n", num2, num3, num2/num3)
			break
		}
	}

}
