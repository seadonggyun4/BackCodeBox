// 당신은 Go 동물원의 사육사!
// 새로 들이고 싶은 동물들을 입력하시고 동물을의 리스트를 출력하세요!
// 동물의 입력은 adopt 함수 에서 해결!  동물 리스트 출력은 truck함수 에서 해결!

package main

import "fmt"

func truck(animals ...string)(count int, list []string){//입력한 동물들을 매개변수로 받고 리스트로 출력한다.
	for i := 0; i < len(animals); i++ {
		list = append(list, animals[i])
		count++
	}
	return
}


func adopt()(animals []string){//동물들을 입력한다.
	for{
		var animal string
		fmt.Println("입양하고 싶은 동물을 입력하세요! :")
		fmt.Scanln(&animal)

		if animal != "stop"{
			animals = append(animals, animal)
		} else {
			fmt.Println("동물 입양이 끝났습니다!")
			break
		}
	}
	return
	
}


func main() {
	fmt.Println("===== Go ZOO 전산시스템 입니다 =====")
	count, list := truck(adopt()...)
	fmt.Printf("입양하려는 동물 수는 %d \n", count)
	fmt.Printf("입양하려는 동물 리스트는 %s \n", list)
}