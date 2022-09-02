package main

import "fmt"

var a int = -1 // int, int8, int16 int32, int64 ->정수(음수포함)
var b uint = 1 // uint, uint8, uint16, uint32, uint64 = uintptr -> 정수(양수,0)
// int, uint 는 컴퓨터가 64비트라면 64 , 32라면 32로 선언된다.

var c float32 = -1.2   // float32, float64 -> 실수형
var d complex64 = -1.2 // complex64, complex128 -> 복소수
//실수와 복소수는 표현범위가 넓은만큼 표현하려는 변수에 알맞는 크기 "정밀도" 가 중요!
// 복소수 = 실수 + 허수

var byt byte = 1  // byte = uint8 -> 정수(양수, 0)
var run rune = -1 // rune = int32 -> 정수(음수포함)

var str string = "hello \n"  // 문자열 타입은 16byte이다.
var str2 string = `hello \n` // 문자열 타입은 16byte이다.

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) // (결과 값: -1.2+0i    =>  실수 + 허수)
	fmt.Println(byt)
	fmt.Println(run)
	fmt.Println(str)
	fmt.Println(str2) // (결과 값:  hello \n  =>  `` 를쓰면 그 안의 내용 그대로 출력된다. )

}
