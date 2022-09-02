package main

import "fmt"

type squre struct{
	width int
	height int
}

func(s squre)squreAreaVal()int{
	s.height += 10
	return s.width * s.height
}

func(s *squre)squreAreaRef()int{
	s.height += 10
	return s.width * s.height
}

func main() {
	squreA := new(squre)
	squreA.width = 10
	squreA.height = 2

	squre_val := squreA.squreAreaVal()
	fmt.Printf("높이:%d X 가로:%d = 넓이:%d \n",squreA.height,squreA.width,squre_val)
	sure_ref := squreA.squreAreaRef()
	//참조가 되면 당연히 원본값도 바뀌기 때문에 올바른 비교를 위해서는 아래쪽에 배치
	fmt.Printf("높이:%d X 가로:%d = 넓이:%d \n",squreA.height,squreA.width,sure_ref)
}