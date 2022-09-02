//[인터페이스]
// 구조체 -> 같은 속성 필드들의 집합체
// 메소드 -> 구조체의 속성을 수행하는 특별한 함수
// 인터페이스 -> 메소드를 모아놓은 집합체

// =============================== 인터페이스를 사용 해서 원, 사각형 구하는 공식! =======================
package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
}

type Rect struct {//사각형 구조체
	width, height float64
}

type Circle struct {//원 구조체
	radius float64
}

func (r Rect) area() float64 {//사각형 메소드
	return r.width * r.height
}

func (c Circle) area() float64 {//원 메소드
	return math.Pi * c.radius * c.radius
}


func main() {
	r1 := Rect{10, 20} //사각형 구조체 변수에 할당하고  초기화
	c1 := Circle{10} // 원 구조체 변수에 할당하고 초기화
	r2 := Rect{12, 14}
	c2 := Circle{5}
/*
	[기존방식]
	fmt.Println(r1.area())
	fmt.Println(c1.area())
	fmt.Println(r2.area())
	fmt.Println(c2.area())
*/

	printMeasure(r1, c1, r2, c2)// 함수 생성 -> 기존 일일이 호출했던 메소드를 하나로 묶어 호출시켜줄 함수
}

func printMeasure(m ...geometry) { //인터페이스를 가변 인자로 하는 함수
	for _, val := range m { //가변 인자 함수의 값은 슬라이스형
		fmt.Println(val.area()) //인터페이스의 메소드 호출
		
		//range문을 순회하며 인터페이스geometry에 들어있던 area메소드 값들을 출력한다.
	}
}