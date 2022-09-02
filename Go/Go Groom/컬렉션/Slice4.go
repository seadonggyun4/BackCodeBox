//[Slice4]

package main

import "fmt"

func main() {

	 A := []int{1,2,3,4,5}
	 B := make([]int, len(A),  cap(A)*2 ) // make()함수 내부에 변수명으로 선언시 cap인지 len 인지 지정을 해줘야함

	 copy(B,A)// copy(붙혀넣을 슬라이스, 복사할 슬라이스)
	 
	 fmt.Println("[B가 A를 copy 했을때]")
	 fmt.Printf("A슬라이스 값:%d  길이:%d  용량:%d \n",A,len(A),cap(A))
	 fmt.Printf("B슬라이스 값:%d  길이:%d  용량:%d \n",B,len(B),cap(B))
	
}