//[구조체 포인터]

package main

import "fmt"

type person struct{
	name string
	age int
	contact string
}

//매개변수로 들어온 값의 person구조체 원래값을 참조해서 +4를 하는 함수
func addAgeRef(a *person){//Pass by reference(참조로 전달)
	a.age += 4// 자동 역참조 -> *생략
}

//매개변수로 들어온 person구조체 값에 +4를 하는 함수
func addAgeVal(a person){//Pass by value(값으로 전달) 
	a.age += 4
	fmt.Println(a)//a.age에 처리된 값은 지역변수로 처리되 addAgeVal 함수가 끝난후 메모리에서 사라진다.
}

func main() {
	/*
	[구조체 포인터 생성 방법 2가지]
	1. "new(구조체이름)"을 사용하여 객체를 생성하기
	2. 구조체 이름 앞에 & 붙이기.
	*/
	p1 := new(person) //포인터 구조체 객체 생성(실제 struct 구조체 메모리를 참조)
	p2 := person{} //빈 구조체 객체 생성

	fmt.Println("포인터 구조체 p1 :",p1)
	fmt.Println("일반 구조체 p2 :",p2)

	//각각 구조체에 age = 25로 초기화
	p1.age = 25
	p2.age = 25

	addAgeRef(p1)//    참조 함수(포인터 구조체) -> 구조체가 포인터라 "&" 안붙혀도 된다.
	addAgeVal(p2)//    값 함수(구조체) -> 구조체 person의 값을 복사해 와서 

	fmt.Println("참조함수 age  출력 :",p1.age) 
	fmt.Println("값전달함수 age 출력 :",p2.age)

	addAgeRef(&p2)//&를 붙이면 "pass by reference"가능
	fmt.Println(p2.age)

}