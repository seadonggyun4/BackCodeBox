//당신은  올해의 GO 중고차 판매왕 입니다.
//우선 매입하고자 하는 중고차 를 입력하시고 리스트로 출력해보세요!

package main

import "fmt"

func carlist(cars []string)(count int, list []string){ //매개변수 값을 (cars ...string) => (cars []string)으로 표현해도 된다.
	for i := 0; i < len(cars); i++ {
		list = append(list, cars[i])
		count++
	}
	return

}

func buycar()(cars []string){
	for{
		var car string
		fmt.Println("매입하려는 차 기종을 입력하세요 :")
		fmt.Scanln(&car)

		if car != "stopbuy"{
			cars = append(cars, car)
		} else {
			fmt.Println("구매종료")
			break
		}
	}
	return
}


func main() {
	fmt.Println("===== GO 중고차 매입 시스템 ====")
	count, list := carlist(buycar())
	fmt.Printf("매입한 중고차 갯수 : %d \n",count)
	fmt.Printf("매입한 중고차 리스트 : %s \n",list)

	
}