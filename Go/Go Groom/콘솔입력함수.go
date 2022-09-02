package main

import "fmt"

func main() {
	var name string
	var age int
	var year int
	var moth int
	var day int

	fmt.Println("주민등록증 발급을 위한 절차 입니다.")
	fmt.Print("1. 이름과 나이를 입력하세요: ")
	fmt.Scan(&name, &age)
	fmt.Println("")

	fmt.Print("2. 태어난 년도, 월, 일 을 순서대로 입력하세요: ")
	fmt.Scanln(&year, &moth, &day)

	fmt.Println("")
	fmt.Println("[주민 등록 증]")
	fmt.Println(name, age)
	fmt.Println(year, "-", moth, "-", day)

}

/*
왜? Scan을 Scanln앞에쓰면 오류가 나는가?

Scanln: 은 숫자를 모두 입력 받고 난후 개행 -> 공백으로 구분하여 개행
Scan: 은 숫자를 입력할때 마다 개행 -> 공백과 개행으로 구분
Scanf:  는 포멧 지정자를 이용여 개발자가  원하는 형태로 맞춰 입력
*/
