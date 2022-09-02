//[function 매개변수!!]
//[3가지의 방법이 있다.]
//1. 값을 복사하는 방법
//2. 값을 참조하는 방법
//3. 가변인자

//[가변인자]
//매개변수의 개수를 고정한 함수가 아니라 함수를 호출할 때마다 매개변수의 개수를 다르게 전달할 수 있도록 만든 함수

//가변인자를 받을때는  "func 함수이름(매개변수이름 ...매개변수형) 반환형"  -> 형 앞에 ....
// 슬라이스를 전달할때는 "함수이름(슬라이스이름...)" -> 슬라이스 이름 뒤에 ...

package main

import "fmt"

func addOne(num ...int)int{//가변인자 사용 -> 배열을 for문으로 선언
	var result int
	for i := 0; i < len(num); i++ {
		result += num[i]
	}
	return result
}


func addTwo(num ...int)int{//가변인자 사용 -> 배열을 range문으로 선언
	var result int
	for _,val := range num{ //for 인덱스, 값 := range 배열 {}
		result += val
	} 
	return result
}


func main() {
	var num1,num2,num3 int = 1,2,3
	var nums []int = []int{10,20,30,40,50}
	fmt.Println(addOne(num1,num2,num3))
	fmt.Println(addOne(nums...))
	fmt.Println(addTwo(num1,num2,num3))
	fmt.Println(addTwo(nums...))
	// 어떤값이 들어가라도 그 값에 맞게 길이가 동적으로 할당된다.
}