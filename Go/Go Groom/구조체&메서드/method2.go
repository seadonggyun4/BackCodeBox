//[메소드:포인트 리시버]
//1. 메소드에 "포인터 정보"를 전달한다면 구조체 필드 값을 메소드에서 직접 접근해 수정할 수 있습니다.
//2. 메소드를 호출할 대는 다른 점이 없지만 메소드의 receiver부분에서 주솟값을 참조하는 연산자인 "*"를 구조체 이름 앞에 붙여주면 됩니다.
package main

import "fmt"

type triangle struct{
	width  int
	height int
}

func (s triangle)triAreaVal() int{//func(구조체 값 정보) -> 메소드를빠져나가면 width,height값 정보 소멸
	s.width += 10
	return s.width *s.height / 2
}

func (s *triangle)triAreaRef()int{//func(구조체 포인터 정보) -> 메소드를 빠져나가도 width,height값 유지
	s.width += 10
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.height = 5
	tri1.width = 10

	triarea_val := tri1.triAreaVal()//값 계산방식
	fmt.Printf("값 계산방식: (%d X %d)/2=%d \n",tri1.width,tri1.height,triarea_val)
	triarea_ref := tri1.triAreaRef()//포인터 계산방식
	fmt.Printf("포인터 계산방식: (%d X %d)/2=%d \n",tri1.width,tri1.height,triarea_ref)
}