package main

import "fmt"

func main() {
	var num01 int
	var num02 int
	var op string

	fmt.Println("첫 번째 숫자를 입력하세요 : ")
	_, err01 := fmt.Scanf("%d\n", &num01)

	if err01 != nil {
		fmt.Println("잘 못 입력했습니다.")

	}

	fmt.Println("두 번째 숫자를 입력하세요: ")

	_, err02 := fmt.Scanf("%d\n", &num02)

	if err02 != nil {
		fmt.Println("잘 못 입력했습니다.")
	}

	fmt.Println("연산할 연산자를 입력하세요(+, -, *, /) : ")

	_, err03 := fmt.Scanf("%d\n", &op)

	if err03 != nil {
		fmt.Println("잘 못 입력했습니다.")
	}

	// if op == "+" {
	// 	fmt.Printf("이 식의 값은 %d", num01+num02)
	// } else if op == "-" {
	// 	fmt.Printf("이 식의 값은 %d", num01-num02)
	// } else if op == "*" {
	// 	fmt.Printf("이 식의 값은 %d", num01*num02)
	// } else if op == "/" {
	// 	fmt.Printf("이 식의 값은 %d", num01/num02)
	// }

	switch op {
	case "+":
		fmt.Printf("이 식의 값은 %d", num01+num02)
	case "-":
		fmt.Printf("이 식의 값은 %d", num01+num02)
	case "*":
		fmt.Printf("이 식의 값은 %d", num01+num02)
	case "/":
		fmt.Printf("이 식의 값은 %d", num01+num02)
	}

}
