//[function 매개변수!!]
//[3가지의 방법이 있다.]
//1. 값을 복사하는 방법
//2. 값을 참조하는 방법
//3. 가변인자

//[값을 복사하는 방법!!]
package main

import "fmt"

func apple(a int){
	a = a * 4
	fmt.Println("함수 apple 에서의 a값 :",a)
}


func main() {
	var a int = 5
	apple(a)//apple함수-> 메인함수의 값을 복사해 apple함수에서만 사용되는 매개변수 생성
	fmt.Println("함수 main에서의 a값 :",a) // 복사된 값은 원본값에 영향을 미치지 못한다.


}