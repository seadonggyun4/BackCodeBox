package main

import "fmt"

func main() {

	//덧셈
	num := 0
	for i := 0; i < 10; i++ {
		num = i + num
	}
	fmt.Printf("%d ", num)

	//별삼각
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			print("*")
		}
		println("")
	}

	println("")
	println("")
	//별삼각 반대
	for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			print("*")
		}
		println("")
	}
	// i가 0일때는 j는 5번, i가 1일때는 j는 4번 이런식으로 줄어든다.

	//while 과 같이 쓰이는 for문
	fmt.Println("[3의배수 숫자 카운트]")
	for i := 1; i < 100; { //조건식이 충족할때(while) 내용 반복!
		fmt.Printf("카운트 :%d \n", i)
		i *= 3
	}

	println("")
	println("")
	//============================================for Range문!!==================================================
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("[배열 arr[6]의  값들]")
	for index, arr := range arr {
		fmt.Printf("index :%d  value: %d \n", index, arr)
	}
	// 값들만 출력 -> 인덱스 변수 자리에 " _ " 넣기
	fmt.Println("")
	for _, arr := range arr {
		fmt.Printf("value: %d\n", arr)
	}
	// 인덱스만 출력 ->  인덱스 변수이름 하나만 선언
	fmt.Println("")
	for index := range arr {
		fmt.Printf("index: %d\n", index)
	}

	// for range 활용시 맵 활용
	var animal map[string]string = map[string]string{
		"dug": "woof",
		"cat": "meow",
		"cow": "moow",
	}

	// map 사용시 인덱스가 꼭 정수가 아니여도 된다.
	fmt.Println("")
	fmt.Println("[동물 울음소리]")
	for animal, bark := range animal {
		fmt.Printf("animal: %s  bark:%s \n", animal, bark)
	}

}

/*
range 키워드는 컬렉션(배열,슬라이스, 맵)으로부터
위치인덱스와 값을 리턴 받아 각각 index와 변수 자리에 할당한다.

즉 for range문은 컬렉션의 모든 요소에 접근해 차례대로 리턴할때 사용
*/
