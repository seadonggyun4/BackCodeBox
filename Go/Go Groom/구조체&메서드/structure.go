// [구조체(structure)]
// 1. 구조체는 "하나 이상의 변수를 묶어서 새로운 자료형을 정의하는 Custom data type"입니다.
// 2. "절차지향" 언어 에서 구조체는 "공통되는 속성과 특징에 따라 쓰인다." => Ex) 사람 struct{이름,성별,나이,전화번호,직업....}
// 3. Go 언어는 "객체지향" 언어지만 "클래스,객체,상속" 개념이 없다.
// 4. Go언어는 클래스를 "구조체로 표현한다"
// 5. Go언어는 기존클래스가 필드와 메소드를 함께 갖는것과 달리 "구조체로 표현한 클래스" 에서는 필드 만을 가지고 메소드는 별도로 분리하여 정의

// 구조체 객체를 생성하려면 "객체이름 := 구조체이름{저장할값}"으로 입력해 선언과 동시에 초기화를 할 수 있다.

package main

import "fmt"

type person struct{
	name string
	age int
	contact string
}


func main() {
	var p1 = person{} // 비어있는 구조체 대입 할당
	fmt.Println(p1)

	//p1 필드 생성후 차례로 값 초기화
	p1.name = "kim"
	p1.age = 25
	p1.contact = "010-2801-4657"
	fmt.Println(p1)

	//p2필드 생성과 동시에값 초기화-> 순서대로 필드에 대입
	p2 := person{"Han", 27, "010-3674-5462"}
	fmt.Println(p2)

	//p3필드 생성과 동시에 값 초기화 -> 이름을 지정해줘 필드에 초기화
	p3 := person{name:"Joe", contact: "010-2742-5325",age: 21}
	fmt.Println(p3)

	//저장된 값 수정
	p1.name = "you"
	fmt.Println(p1)
	//특정값 에만 접근
	fmt.Println(p1.name)

}