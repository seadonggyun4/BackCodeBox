package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i] //arr배열 값을 차례로 clone배열에 복사
	}
	fmt.Println(clone)
}
