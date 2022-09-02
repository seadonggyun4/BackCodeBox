// 배열의 역순 !!!!!!!!!!!!!!

package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}

	for i := 0; i < len(arr); i++ {
		temp[i] = arr[(len(arr)-1)-i] // 저장공간은 5 주솟값은 0~4이기 때문에 맨 뒤 공간인 4부터 참조
	}

	for i := 0; i < len(arr); i++ { //역순으로된 temp를 arr에 대입
		arr[i] = temp[i]
		// 굳이 len으로 배열 주소길이를 계산하는 이유는 배열의 크기가 명확하지 않거나 변경될수 도 있기때문
	}
	fmt.Println("temp : ", temp)
	fmt.Println("arr : ", arr)
}
