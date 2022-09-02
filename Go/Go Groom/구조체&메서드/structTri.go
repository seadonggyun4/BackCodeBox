package main

import "fmt"

type triangle struct{
	width int
	height  int
}

func triArea(s *triangle)int{
	return s.height * s.width /2
}

func main() {
	tri1 := new(triangle) // "new"로 생성한 구조체 객체는 포인터값 반환
	tri1.width = 10
	tri1.height = 20
	triArea := triArea(tri1)//바로 triangle 구조체를 가져올수는 없다 이를 담고있는 변수는 반드시 존재해야함

	fmt.Printf("삼각형 높이:%d, 너비:%d 일때 넓이:%d",tri1.height,tri1.width,triArea)
}