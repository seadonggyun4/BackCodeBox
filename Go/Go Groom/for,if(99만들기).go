package main

import "fmt"

func main() {
	fmt.Println("[=====두 수를 더해서 99 만들기!=====]")
	// var A int
	// var B int
	var count int =0 

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {			
			count += 1 // 너무나 당연하게도 반복횟수를 새고싶을땐 1씩 더하면 된다.
			if ((i*10+j)+(i + j*10)) == 99 {
				
				fmt.Printf("%02d + %02d = %d \n" , i*10 + j ,i + j*10, (i*10+j)+(i + j*10))
				//   %02d -> 정수형 앞에 0을 붙혀 출력해준다.
			}
		}
		
	}
	fmt.Println("반복횟수:",count)

}


/*
package main

import "fmt"

func main() {
	
	for i := 0; i < 10; i++ {
		for j:= 0; j < 10; j++{
			if (i == j) || (i + j != 9) {
				continue
			}
			
			fmt.Printf("%02d + %02d = %02d\n", i*10 + j, j*10 + i, (i + j)*10 + i + j)
		}
	}
}

*/