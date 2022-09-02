//[인터페이스:빈 인터페이스]
// Go 언어에서 함수는 type처럼 쓸수있다 -> 익명함수:일급함수
// Go 언어에서 인터페이스도 type처럼 쓸수 있다.

// [특징]
// 1.인터페이스는 내용을 따로 선언하지 않아도 형으로서 사용할 수 있습니다.
// 2.인터페이스는 하나의 형이기 때문에 매개변수로 사용될 수 있습니다.
// 3.인터페이스는 어떠한 타입도 담을수 있는 컨테이너입니다. 즉 "Dynamic type"입니다.

//=> 빈 인터페이스 형을 쓰면 어떠한 형을 쓰던 편하게 담을수 있다.

package main

import "fmt"

func printVal(i interface{}){//인터페이스를 매개변수로 사용하고 있다.
	fmt.Println(i)
}

func main() {
	var x interface{}//빈 인터페이스 생성 -> 메소드를 담기 위함이 아니라 Dynamic Type으로서 사용하기 위함!!!!!

	x = 1 //int type을 인터페이스에 넣었다
	printVal(x)

	x = "test"// string type을 인터페이스에 넣었다.
	printVal(x)
	
}