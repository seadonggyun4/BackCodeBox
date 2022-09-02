package main

import "fmt"

func main (){
	// fmt.Println("switch study")
	// fmt.Println("switch문은 조건에 따른 분기문이다.")
	// fmt.Println("if, else if 문은 조건과 맞는지 하나씩 검사하는 느낌")
	// fmt.Println("switch문은 case 라벨이 맞는 곳을 찾아 내어 바로 그 구문의 조건문을 실행")
	fmt.Println("[World of RPG]")
	fmt.Println("캐릭터를 고르세요(직업입력!) : ")
	fmt.Println("1.마법사 2.전사 3.궁수 4.암살자 ")

	var char string
	fmt.Scan(&char)


	if char == "마법사"{
		fmt.Println("======================")
		fmt.Println("마법사를 고르셨습니다.")
		fmt.Println("======================")
	} else if char == "전사"{
		fmt.Println("======================")
		fmt.Println("전사를 고르셨습니다.")
		fmt.Println("======================")
	} else if char == "궁수"{
		fmt.Println("======================")
		fmt.Println("궁수를 고르셨습니다.")
		fmt.Println("======================")
	} else if char == "암살자"{
		fmt.Println("======================")
		fmt.Println("암살자를 고르셨습니다.")
		fmt.Println("======================")
	} else {
		fmt.Println("잘못고르셨습니다. 다시선택하세요")
		
	} 	


	switch{
	case char == "마법사" :
		fmt.Println("HP:50")
		fmt.Println("MP:150")
		fmt.Println("Speed:10")
	case char == "전사" :
		fmt.Println("HP:150")
		fmt.Println("MP:50")
		fmt.Println("Speed:10")
	case char == "궁수" :
		fmt.Println("HP:100")
		fmt.Println("MP:100")
		fmt.Println("Speed:20")
	case char == "암살자" :
		fmt.Println("HP:80")
		fmt.Println("MP:80")
		fmt.Println("Speed:80")
	}

	if char == "궁수" || char == "전사" || char == "마법사" || char == "암살자"{
		fmt.Println("")
		fmt.Println("~~~~~~~~~~~~~~~~~~~ 이제 모험을 떠나세요 ~~~~~~~~~~~~~~~~~~~")
	}
	
	
}

/*
switch 태그/표현식{
case 라벨/표현식:
	수행구문
case 라벨/표현식:
	수행구문
	...
default 라벨/표현식:
	수행구문
}

*/