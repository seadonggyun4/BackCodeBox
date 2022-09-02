package main

import "fmt"

//함수 type 지정
type calculatorNum func(a int, b int)int
type calculatorStr func(a string,b string)string


func calNum(f calculatorNum,a int, b int)int{
	result := f(a,b)
	return result
}

func calStr(f calculatorStr,a string, b string)string{
	result := f(a,b)
	return result
}

func main() {
	Pluse := func(a int, b int)int{return a + b}
	Minuse := func(a int, b int)int{return a - b}
	PluseStr := func(a string, b string)string{return a + b}

	calcA := calNum(Pluse, 10, 20)
	fmt.Println(calcA)

	clacB := calNum(Minuse,10,20)
	fmt.Println(clacB)

	calcC := calStr(PluseStr,"Hello","World")
	fmt.Println(calcC)


}