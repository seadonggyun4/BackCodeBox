//[메소드]
//특정 속성들(구조체)을 이용해 기능을 수행하기 위해 만들어진 특별한 함수 "메소드"

//[메소드 선언방법]
//1. func(매개변수이름 구조체이름)메소드이름()반환형{} -> Ex)  func(s triangle)triArea()float32{}
//2. "함수이름"을 입력하지 않고 구조체이름 뒤에 메소드 이름을 입력합니다. 본문에서 메소드를 이용하기 위해 이름을 사용합니다.

package main

import "fmt"

type triangle struct{
	width int
	height int
}

// triArea는 메소드
// s triangle은 "receiver" -> 어떤 구조체를 전달받는지 명시. 
//구조체 객체 자체를 전달받는것이아니라 구조체 객체 정보를 전달 받고 메소드의 기능을 수행하는것.
//함수를 사용해서 매개변수로서 객체를 활용하는 모습과는 차이가 있다.
func (s triangle)triArea()int{
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12
	tri1.height = 5

	triArea := tri1.triArea()

	fmt.Printf("삼각형 길이:%d  X  높이:%d = 넓이: %d",tri1.width,tri1.height,triArea)
}