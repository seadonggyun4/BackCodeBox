//[Slice 부분 복사]
//copy 이외에도 부분복사가 가능하다.

// 변수 := 슬라이스[시작인덱스 : 끝인덱스+1] or 변수 := 슬라이스[시작인덱스 : 배열 전체 용량]

package main

import "fmt"
func main() {
	a := [5]int{}
	var b []int
	for i := 0; i < cap(a); i++ {
	 	b = append(b,i) // a 인덱스 값 저장
	}
	fmt.Println("b슬라이스 값 :",b)

	c := b[0:3] // 변수 := 슬라이스[시작인덱스 : 끝인덱스+1] 
	fmt.Println("b [0:3]부분 복사한 c 값 :",c)

	c[0] = 5
	fmt.Println("b [0:3]부분 복사한 c 값에 0번 인덱스에 5를 대입 :",c)

	fmt.Println("돌고돌아 b 슬라이스 값",b,"\n","\n","\n")

	fmt.Println("=======================   슬라이스는 값을 복사하는게 아닌 참조한다!!! ====================== \n\n")


}

// 복사해온 슬라이스의 값을 바꾸면 기존 복사한 슬라이스의 값도 바뀐 것을 확인할 수 있다.
// 왜냐하면 슬라이스는 배열과 다르게 값을 복사해오는 것이 아니라
//  슬라이스 자체가 참조하고있는 주소값을 같이 참조하는 것을 의미하기 때문입니다.
// 하지만 같은 상황이라면 배열은 단순히 값을 복사해서 초기화합니다