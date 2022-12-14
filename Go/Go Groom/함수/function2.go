// [function 지역변수 전역변수]
// 매개변수는 값 자체를 전달하는 방식, 값의 주소를 전달하는 방식 두가지가있다.(값을 복사해오는 방식과, 값을 참조하는방식)
// 지역변수= 특정함수 내부에 선언된 변수,  전역변수= 함수외부에 선언된 변수

// 지역변수 전역변수 차이점
// 1. 메모리에 존재하는 시간
// - 지역변수는 해당 지역에서 선언되는 순간 메모리가 생성, 벗어나면 자동 소멸
// - 전역변수는 코드가 시작되어 선언되는 순간 메모리생성되고 코드 전체가 끝날때까지 메모리를 차지
package main

import "fmt"
var a int = 1 //전역변수

func global()int{
	return a+ 1
}


func local()int{
	a := 50//전역으로 선언된 a 위에 겹쳐지게 된다. local함수가 끝나면 사라지게 된다.
	return a
}


func main() {

	fmt.Println("전역변수 a 사용 :",global())//전역변수 사용
	fmt.Println("지역변수 a 사용 :",local())//지역변수 사용

	
}

// 전역변수의 선언은 가급적 피해야합니다.  전역변수는 프로그램의 구조를 복잡하게 만들고, 프로그램이 끝날때까지 메모리를 차지하고있기 때문입니다. 
// 따라서 전역변수를 사용하는 것은 신중해야합니다.