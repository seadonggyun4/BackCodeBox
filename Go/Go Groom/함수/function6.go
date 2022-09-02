// [반환값(리턴값)]
// 반환값은 함수의 기본적인 기능중 하나인 연산에대한 출력을 하기위한 결과 값
// Go언어에서는 복수개의 반환값을 반환할 수 있다는 특징이 있다.

// 1. 반환값의 개수만큼 반환형을 명시해야한다. 2개 이상의 반환형을 입력할 때는 괄호(())안에 명시
// 2. 동일한 반환형이더라도 모두 명시해야한다.((int,int,int))

package main

import "fmt"

func addOne (num ...int)(int, int){//리턴값에 인트는 반드시 선언해줘야함
	var result int
	var count int

	for i := 0; i < len(num); i++ {
		result += num[i]
		count++
	}
	return result, count//리턴값 두개
}


func main() {
	var nums []int = []int{10,20,30,40}
	fmt.Println(addOne(nums...))
}