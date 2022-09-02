package main

import "fmt"


func calc(f func(int,int)int, a int, b int)int{
	result := f(a,b)
	return result
}
/*calc 함수는 인자로 빈껍대기 함수 f, 변수a,b를 받아서
f함수와 변수 a,b를 조립하는 함수 ->calc을 만든것이다.
*/


func main() {
	Pluse := func(a int, b int)int{return a + b}
	Minuse := func(a int, b int)int{return a - b}
	
	calcA := calc(Pluse, 10, 20)
	fmt.Println("clac 매개변수 Pluse함수 :",calcA)

	calcB := calc(Minuse,10,20)
	fmt.Println("clac 매개변수 Minus함수 :",calcB)
}