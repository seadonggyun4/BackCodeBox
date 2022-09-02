package main

import "fmt"

func main() {
	s := "Hello World"

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ",") // 문자열을 rune 타입으로 변환후 다시 문자열로 변환
	}
}
