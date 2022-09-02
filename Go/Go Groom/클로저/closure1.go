//[클로저(외부변수 접근)]
/*클로저는 함수 안에서 "익명함수"를 정의해서 바깥쪽 함수에 선언한 변수에도
접근할 수 있는 함수*/

package main

import "fmt"


func next() (func()int){ //next함수 : 매개변수(),  반환값(func()int)
	i := 0
	return func()int{
		i += 1 
		return i
	}
}
	/*
	출력결과를 확인시 nextInt를 실행할때마다 초기화 되는것이 아니라 계속해서  값이 증가한다.
	왜냐하면 i의 연산 기능을 하는 익명 함수 안에서 i가 선언되지 않고 익명 함수 밖에 있는 변수 i를 참조하고있기 때문
	익명함수 자체가 지역변수로 i를 갖는 것이 아니기 때문에 외부 변수 i가 상태를 계속 유지하면서
	값을 1씩 증가시키는 기능을 하게된다.	
	
	=> 익명함수가 바깥쪽에 선언된 변수에 접근 가능 [클로저]
	*/


func main() {
	nextInt := next()//변수에 next()함수 할당

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := next()//변수에 next()함수 할당
	fmt.Println(newInt())
}