//[Slice3]

package main

import "fmt"

func main() {
	SliceA := []int{1,2,3} //1,2,3 값을 참조하는 슬라이스A생성
	SliceB := []int{4,5,6} //4,5,6 값을 참조하는 슬라이스B생성

	SliceA = append(SliceA, SliceB...) // append(슬라이스A, SliceB 요소들)
	fmt.Println(SliceA)
}