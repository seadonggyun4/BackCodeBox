package main

import "fmt"

func main (){

first:	
	fmt.Println("[Go-Online RPG]")
	fmt.Println("캐릭터를 고르세요(직업입력!) : ")
	fmt.Println("1.마법사 2.전사 3.궁수 4.암살자 ")

	Loding := 150 // 로딩
	var char string//캐릭터
	fmt.Scan(&char)// 클래스 선택

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
		fmt.Println("캐릭터를 잘못고르셨습니다. 다시선택하세요")
		goto first
	} 	


	switch{// 캐릭터 능력치
	case char == "마법사" :
		fmt.Println("HP:50")
		fmt.Println("MP:150")
		fmt.Println("Speed:10")
		fmt.Println("======================")
	case char == "전사" :
		fmt.Println("HP:150")
		fmt.Println("MP:50")
		fmt.Println("Speed:10")
		fmt.Println("======================")
	case char == "궁수" :
		fmt.Println("HP:100")
		fmt.Println("MP:100")
		fmt.Println("Speed:20")
		fmt.Println("======================")
	case char == "암살자" :
		fmt.Println("HP:80")
		fmt.Println("MP:80")
		fmt.Println("Speed:80")
		fmt.Println("======================")
	}

	var classchoice int
	fmt.Println("여러분의 직업이 마음에 드시나요?")
	fmt.Println("다음 두 선택지중 골라주세요 !")
	fmt.Println("1:네!                 2:아니요! ")
	fmt.Scan(&classchoice)//캐릭터 선택이 마음에 드는지?
	
	if classchoice ==  1{//시나리오1:  캐릭터 선택이 마음에 든다.
		
		if char == "궁수" || char == "전사" || char == "마법사" || char == "암살자"{
			fmt.Println("")
			fmt.Println("~~~~~~~~~~~~~~~~~~~ 이제 모험을 떠나세요 ~~~~~~~~~~~~~~~~~~~")
			fmt.Println("~~~~~~~~ 여러분은 이제 시작의 마을로 가게 될것입니다!!!~~~~~~~~")
			fmt.Println("~~~~~~~~ 꾸준히 성장하셔서 버그 마왕을 무찔러 주세요!~~~~~~~~")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~ [LODING] ~~~~~~~~~~~~~~~~~~~~~~~")
	
			for i := 0; i < Loding; i++ {
				i := "="
				print(i)
			}
			print("\n")
		}
	} else{ // 캐릭터 선택이 마음에 안든다.
		fmt.Println("")
			fmt.Println("~~~~~~~~~~~~~~~~~~~ 저런 캐릭터를 잘못 고르셨군요! ~~~~~~~~~~~~~~~~~~~")
			fmt.Println("~~~~~~~~ 걱정하지 마세요 캐릭터는 다시 고르시면 됩니다!!!~~~~~~~~")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~ [LODING] ~~~~~~~~~~~~~~~~~~~~~~~")
			for i := 0; i < Loding; i++ {
				i := "="
				print(i)
			}
			print("\n")
		goto first
	}



	
	
	}