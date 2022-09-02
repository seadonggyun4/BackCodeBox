//[map 내부 값 존재 여부 확인]

package main

import "fmt"

func main() {
	m := map[string]string{
		"apple":"red",
		"banana": "yellow",
		"grape" : "purple",
	}
	fmt.Println(m)
	fmt.Println(m["apple"])
	fmt.Println(m["mago"])// 존재하지 않는 인덱스 이기 때문에 ""  같은 빈칸으로 출력
	
	val, exist := m["apple"] // val은 인덱스에 해당하는 값 , exist는 존재여부를 bool타입으로 받는다. -> 변수 이름은 상관 없다.
	fmt.Println(val, exist) 
	val,exist = m["mago"]
	fmt.Println(val,exist)

	val = m["apple"] // 값만 받고싶을때는 변수 하나만 선언
	fmt.Println(val)

	_, exist = m["apple"]// 값의 존재 여부만 알고싶을때는 (_,변수) 와 같이 선언
	fmt.Println(exist)
}