// [익명 함수]

// [익명함수가 존재하게 된 이유]
// 1. 함수 선언 자체가 프로그래밍 전역으로 초기화 되면서 메모리를 잡아먹는다.
// 2. 함수가 가지고 있는 기능을 수행할 때마다 함수를 찾아서 호출해야 하게 된다.

// [익명함수의 특징]
// 1. 함수의 이름만 없고 그 외에 형태는 동일하다.
// 2. 함수의 블록 마지막에 괄호를 사용해 매개변수를 넣고 바로 호출할수 있다.
// => 익명 함수는 함수의 "기능적인 요소"만 어디서든 가볍게 활용하기 위해 사용한는 것입니다.

package main

import "fmt"

func addDeclared(nums ...int)(result int){//이름이 있는 선언함수 -> 프로그램 시작과 동시에 모두 읽는다.
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

func main() {
	var nums []int= []int{10,5,5,10}
	

	//이름이 없는 익명함수
	// 변수에 초기화한 함수는 변수 이름을 함수의 이름처럼 사용할 수 있다. -> 해당함수가 실행되는 동시에 읽는다.
	addAnonymous := func(nums ...int)(result int){
		for i := 0; i < len(nums); i++ {
			result +=nums[i]
		}
		return
	}

	fmt.Println(addDeclared(nums...))
	fmt.Println(addAnonymous(nums...))
}