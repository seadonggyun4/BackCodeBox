// 배열(Array//
package main

import "fmt"

func main() {
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Printf("정석 출력 :%d \n",a)
	b := [5]int{1,2,3,4,5}
	fmt.Printf("약식 출력 :%d \n", b)
	c := [...]int{1,2,3,4,5}
	fmt.Printf(".... 출력 :%d \n",c)

}