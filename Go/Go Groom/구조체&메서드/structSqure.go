package main

import "fmt"


type Squre struct {
	weight int
	height int
}

func SqureArea(s *Squre) int {
	return s.height * s.weight
}

func main() {
	Squre1 := Squre{} // 포인터 형이 아닌 일반 형으로 받음
	Squre1.height = 20
	Squre1.weight = 15
	SqureArea := SqureArea(&Squre1) // 계산 함수에 전달해 줄때 포인터 형으로 매개변수를 전달
	fmt.Println(SqureArea)

	//[new를 써서 포인터 형으로 받는다.]
/*
	Squre2 := new(Squre)
	Squre2.height = 30
	Squre2.weight = 20
	SqureArea := SqureArea(Squre2)
	fmt.Println(SqureArea)
*/

}