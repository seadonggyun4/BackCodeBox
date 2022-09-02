// [반환값(리턴값)]
// 반환값은 함수의 기본적인 기능중 하나인 연산에대한 출력을 하기위한 결과 값
// Go언어에서는 복수개의 반환값을 반환할 수 있다는 특징이 있다.

//Named retrun parameter 는 '이름이 붙여진 반환인자'입니다.
//1. (반환값이름1 반환형1, 반환값이름2 반환형2,...)형식으로 입력합니다.
//2. (반환값이름 반환형) 자체가 인수 선언 입니다.
// EX) (count int, list []string)

// 함수를 변수처럼 사용할수 있다.
// 함수를 변수처럼 담는 행위를 너무 남발하면 코드가 복잡하고 지저분해 질 수도 있다.
package main

import "fmt"

func dish(foods ...string)(count int, putdish []string){ //order의 리턴값을 매개변수로 받아 연산을 진행한다.
	
	for i := 0; i < len(foods); i++ {
		putdish = append(putdish, foods[i])
		count ++
	}
	return
}

func order()(foods []string){// 변수의 역할을 하게될 함수

	for{
		var food string
		fmt.Println("음식을 입력받으세요 :")
		fmt.Scanln(&food)

		if food != "exit"{
			foods = append(foods, food)
		} else {
			fmt.Println("주문이 끝났습니다.")
			break
		}
	}
	return
}


func main() {
	fmt.Println("====== Go 식당에 어서오세요! 메뉴를 골라주십시오 =======")
	count, foods := dish(order()...)// count,foods 값을 담기 위한 함수 생성
	fmt.Printf("주문하신 메뉴 숫자는 %d \n", count)
	fmt.Printf("주문하신 멘뉴 리스트는 %s \n", foods)
	
}