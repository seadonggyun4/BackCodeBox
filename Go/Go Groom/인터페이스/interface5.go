// [인터페이스 : type Assertion]
// 1. 인터페이스 형으로 선언된 변수는 초기화하는 값에 따라 형이 자동명시 되지만 사실상 Danamic type입니다.
// 2. 따라서 확실한 형을 표현하기 위새서 "type Assertion"을 할 필요가 있습니다.
// 3. 인터페이스 형은 nil 값일수 없습니다.
// => "변수이름.(형)" 을 명시하면 됩니다.

package main

import "fmt"

func main() {
	var num interface{} = 10 //인터페이스 선언 및 초기화(int type)

	a := num //int형으로 들어가 있지만 아직 danamic type -> 언제든지 다른 type으로 변환될수 있는상태
	b := num.(int) //int type으로 type assertion을 했다.

	fmt.Printf("%T,%d \n",a,a)
	printtest(b)
}

func printtest(i interface{}){
	fmt.Printf("%T,%d \n",i,i)
}
/*
type assertion(타입표명)
프로그래머가 컴파일러에게 내가 너보다 타입에 더 잘알고 있고, 나의 주장에 대해 의심하지 말라고 표기하는 것 이다.
인터페이스가 가지고 있는 실제값에 접근할 수 있게 해준다.
*/