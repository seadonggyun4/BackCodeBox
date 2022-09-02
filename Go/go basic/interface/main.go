package main

import (
	"fmt"
)

type StructA struct {
	val string
}

// string() -> 코드간 연결 주체
func (s *StructA) String() string {
	return "val:" + s.val
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a)
}

// interface -> 객체간 상호관계 정의
