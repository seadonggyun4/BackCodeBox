//[map 갱신, 삭제, 업데이트]
package main

import "fmt"

func main() {

	var m = make(map[string]string)

	m["Hi"] = "gildong" // 맵[인덱스] = 값
	m["welcom"] = "Go town"
	m["bye"] = "gildong"

	//궁금한점! 이렇게 선언했을때 m에 대입되는 순서는 어떻게 되는지?

	fmt.Println(m)

	m["Hi"] = "donggyun" //"Hi" 인덱스의 값을 업데이트
	fmt.Println(m)

	delete(m, "bye")// m 맵의 "bye"인덱스와 그 값을 제거
	fmt.Println(m)

	
}