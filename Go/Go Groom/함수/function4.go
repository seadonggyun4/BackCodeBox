//[function 매개변수!!]
//[3가지의 방법이 있다.]
//1. 값을 복사하는 방법
//2. 값을 참조하는 방법
//3. 가변인자

//[값을 참조하는 방법!!]
//값을 참조하려면 "포인터를 이용해야 한다."
package main

import "fmt"


func banana(a *int){
	*a =  *a +5
	fmt.Printf("banana함수에 a주소%d 를 넣어 포인터의 원본값으로 받는 식으로 참조한다. \n",a)
	fmt.Printf("banana함수의 변환된 값 :%d \n",*a)
}

func main() {
var a int = 5
banana(&a)
fmt.Printf("원본 a값은 5에서 %d로 변환되었다.",a)
}