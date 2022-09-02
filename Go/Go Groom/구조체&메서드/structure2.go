//[구조체: 생성자 함수]
// 구초체를 사용하기전에 초기화 되어야하는경우 가 있다.
//-> EX)구조체가 map 형식일때 초기화 할때마다 맵 필드도 같이 초기화해야하는 번거로움 있다

// => 구조체 생성과 동시에 초기화 하면 된다.
// => 생성자 함수로 해소

package main

import "fmt"

type mapStruct struct{ //mapStruct 구조체 생성
	data map[int]string // map 타입 "data" 컬럼 생성
}

func newStruct() *mapStruct{ //생성자 함수 newStruct() -> 반환값 mapStruct의 포인터
	d := mapStruct{}//구조체 객체를 생성하고 초기화함
	d.data = map[int]string{}//data 필드의 맵을 초기화함
	return &d// mapStruct를 담고있는 변수d는 "포인터"
}

func main() {
	s1 := newStruct()//생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}//생성자 없이 구조체 호출
	s2.data = map[int]string{}// data 필드의 맵을 초기화 -> 생성자 없이 할 경우 계속해서 호출해줘야 한다.
	s2.data[1] = "hello"
	s2.data[10] = "world"
	fmt.Println(s2)
}
/*
맵 자료형은 포인터를 값으로 갖는 reference타입 자료형이다.
변수의 크기는 포인터의 크기만큼을 갖게되며, 이후 초기화 및 값 할당에 따라 맵의 자료구조의 크기가 정해지고 변하게 된다
구조체 객체의 맵 자료형 필드 역시 마찬가지이다.
*/