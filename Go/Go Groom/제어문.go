package main

import "fmt"

func main() {
	fmt.Println("[Break, Goto,  countinue문 연습]")
	var count int
	var count2 int
	var count3 int

	fmt.Println("[!!!!모든 조건문은 5로 설정!!!!]")

	for i := 0; i < 100; i++ {
		if i ==5{
			continue //조건이 충족되면 반복문 처음부분으로 돌아간다! -> 사용되어있던 if문 은 그대로 지나친다.
		}
		count += 1		
	}

	fmt.Printf("%d 만큼 반복 후 continue로 반복문 탈출!!!\n", count)

	for i := 0; i < 100; i++ {
		if i ==5{
			break//조건이 충족되면 반복문을 깨고 탈출한다.
		}
		count2 += 1
	}
	fmt.Printf("%d 만큼 반복 후 break로 반복문 탈출!!!\n", count2)

// ONE: -> 만약 이 위치에 라벨이 존재시 반복문은 끊이지 않고 계속해서 이어진다. 이점이 continue와 다른점
	for i := 0; i < 100; i++ {
		if i == 5{
			goto ONE // 라벨을 지정해주고 해당 라벨쪽으로 이동을 한다! 프로그램의 흐름을끊는 기법이라 요즘은 사용을 자주 안함
		}
		count3 +=1
	}

	ONE:
	fmt.Printf("%d 만큼 반복후 goto로 반복문 탈출!!!\n", count3)
	


}