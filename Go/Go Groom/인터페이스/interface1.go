// [인터페이스]
// 구조체 -> 같은 속성 필드들의 집합체
// 메소드 -> 구조체의 속성을 수행하는 특별한 함수
// 인터페이스 -> 메소드를 모아놓은 집합체

// ========================= 인터페이스를 사용하지 않고 원, 사각형 구하는 공식! =====================
package main

import (
	"fmt"
	"math"
)

type Rect struct{//사각형 구조체
	width float64
	height float64
}

type Circle struct {//원 구조체
	radius float64
}

func (r Rect)area() float64{//사각형 넓이 연산 메소드
	return r.width * r.height
}

func(c Circle)area() float64{//원 넓이 연산 메소드
	return math.Pi * c.radius * c.radius
}

func main() {
	r1 := Rect{10,20}//사각형 구조체 초기화
	c1 := Circle{10}//원 구조체 초기화
	r2 := Rect{20,30}

	// 전달받는 구조체가 다르기때문에 메소드의 이름이 동일하게 선언 되어도 괜찮습니다.
	fmt.Println(r1.area())
	fmt.Println(c1.area())
	fmt.Println(r2.area())

}