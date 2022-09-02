// 배열의 역순 !!!!!!!!!!!!!!

package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[(len(arr)-1)-i] = arr[(len(arr)-1)-i], arr[i]
		// 이중대입
		// 배열의 주소가 arr[0], arr[4] 일때 = arr[4], arr[0]의 값을 대입
		// arr[i] = arr[(len(arr)-1)-i]이면 마지막 배열의 값이 arr[(len(arr)-1)-i] 에 들어갔을때 -1 이되기 때문에 안된다.
	}
	fmt.Println(arr)
}
