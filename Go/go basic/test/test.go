package main

import "fmt"

func main() {
	var c float32 = 0.1
	var d float32 = 0.2

	fmt.Println(1 == 1)             // true: 두 정수가 같으므로 true
	fmt.Println(3.5 == 3.5)         // true: 두 실수가 같으므로 true
	fmt.Println("Hello" == "Hello") // true: 두 문자열이 같으므로 true
	fmt.Println(c == d)
	fmt.Println(c == d)
	fmt.Println(c == d)

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b) // true: 두 배열이 같으므로 true

	f := []int{1, 2, 3}
	fmt.Println(f == nil) // false: 슬라이스를 nil과 비교하여
	// 메모리가 할당되었는지 확인

	e := map[string]int{"Hello": 1}
	fmt.Println(e == nil) // false: 맵을 nil과 비교하여
	// 메모리가 할당되었는지 확인

	g := 1
	var p1 *int = &g
	var p2 *int = &g
	fmt.Println(p1 == p2) // true: 포인터에 저장된 메모리 주소가 같으므로 true
}
