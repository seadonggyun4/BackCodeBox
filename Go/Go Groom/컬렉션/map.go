//[map]
//배열과 슬라이스는 인덱스번호를 임의로 바꿀수 없지만 맵은 가능하다!
// map은 "key : value" 형식으로 값을 매핑해서 저장합니다. 따라서, 맵을 활용한다면 "love:사랑, root:뿌리, submit:제출하다"와 같은 형식으로 값을 저장할 수 있다.
//슬라이스와 맵의 공통점은 두 컬랙션 모두 값을 직접적으로 저장하는 것이 아닌 '참조 타입(Reference type)'이라는 점입다.

//선언방법 -> " var 변수 map[key자료형]value자료형"

package main

import "fmt"

func main() {
	var a map[int]int
	fmt.Print("a map:")
	fmt.Println(a)
	if a == nil {
		fmt.Println("a map은 nil이다.")
	}

	var b map[string]string = map[string]string{
		"apple": "red",
		"banana": "yellow",
		"grape" : "purple",
	}

	fmt.Println(b)
	fmt.Printf("b map 의 len:%d ",len(b))



}
	
