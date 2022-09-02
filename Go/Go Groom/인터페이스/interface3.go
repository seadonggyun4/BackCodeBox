package main

import (
	"fmt"
	"math"
)

type geometry interface{ 
	area() float64 //넓이 계산 메소드
	perimeter() float64 //둘레 측정 메소드
}

type Rect struct{ // 사각형 구조체
	width float64
	height float64
}

type Circle struct{//원 구조체
	radius float64
}

func (r Rect)area()float64{//사각형 넓이계산 메소드
	return r.width * r.height
}

func (c Circle)area()float64{//원 넓이계산 메소드
	return math.Pi * c.radius * c.radius
}

func(r Rect)perimeter()float64{//사각형 둘레측정 메소드
	return 2 * (r.width + r.height)
}

func(c Circle)perimeter()float64{//원 둘레측정 메소드
	return 2 * math.Pi * c.radius
}

func main() {
	

	r1 := Rect{10,20}
	c1 := Circle{10}
	r2 := Rect{20,30}
	c2 := Circle{30}

	printMeasure(r1,c1,r2,c2)

}

func printMeasure(m ...geometry){
	for _, val := range m {
		fmt.Println("기존값 :",val)
		fmt.Println("넓이 :", val.area())
		fmt.Println("둘레 :", val.perimeter(),"\n\n")
	}
}


/*
1. 인터페이스는 메소드들을 통합적으로 관리할수 있게 해주는 집합이다.
2. 구조체가 늘어났을때 "메소드" 자체를 하나로 통합할 수는 없지만 늘어나는 메소드들을 추적,관리 할수 있게 해 준다.
*/