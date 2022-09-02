//[클로저]
/*클로저는 함수 안에서 "익명함수"를 정의해서 바깥쪽 함수에 선언한 변수에도
접근할 수 있는 함수*/

package main

import "fmt"

func main() {
	a, b := 10,20
	str := "Hello Goorm!"

	result := func () int{//익명함수
		return a + b //main 함수 변수 바로 접근
	}()

	func(){
		fmt.Println(str)//main함수 변수 바로접근
	}()

	fmt.Println(result)
	
}
//1. main함수 내에 선언된 익명함수들이 main()함수의 변수를 매개변수 없이 접근한다.
//2. 함수 안에서 함수를 반환하는 것이 "클로저"