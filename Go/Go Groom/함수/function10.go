//======================================[함수 원형 정의]==============================================

// 함수 원형을 정의해 두면 "일급함수"를 작성할때 매개변수와 반환형을 매번 선언하는 불편함을 극복할 수 있다.

package main

import "fmt"

//type을 사용한 함수 원형 정의
type calculatorNum func(int,int)int
type calculatorStr func(string,string)string

func calNum(f calculatorNum,a int, b int)int{
	//원형정의된 함수를 매개변수로 받음으로써 가독성이 생기고 깔끔해졌다.
	result := f(a,b)
	return result
}

func calStr(f calculatorStr,a string,b string)string{
	sentence := f(a,b)
	return sentence
}

func main() {
	multi := func(a int,b int)int{return a + b}
	duple := func(a string,b string)string{return a + b}

	r1 := calNum(multi,10,20)
	fmt.Println(r1)

	r2 := calStr(duple,"Hello","Go-lang")
	fmt.Println(r2)

}