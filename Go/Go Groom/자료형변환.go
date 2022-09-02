/*
자료형 변환의 종류 = 자동 형 변환(묵시적 형 변환), 강제 형 변환(명시적 형 변환)

Go 언어 에서의 형변환
=> GO언어는 자동 형변환이 안된다!
=> Go언어는 강제 형변환 즉 형변환을 할때 명시를 해줘야만 변환이 가능하다.
*/

package main

import "fmt"

func main() {

	var a int = 5
	var str string = "Hell"
	var c float32 = 3.12

	var b = float32(a)
	// var str_int = []int(str)
	var str_int = []byte(str)
	var d = int(c)

	fmt.Println(b)       //a 변환
	fmt.Println(str_int) // str변환
	fmt.Println(d)       // c 변환

}

//strconvert -> 문자열을 강제로 int형으로 바꿀수 있지만 문자열을 애당초 byte 터압
// byte타입과 숫자인 int 는 아예 다른 성질?

// string 타입은 int(정수로 변환할수 없다.) 오직 byte 타입만 가능?
