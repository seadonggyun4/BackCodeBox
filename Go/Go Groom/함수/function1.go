// [Go 언어 함수]
// 함수 선언 방법 =>   func 함수이름(매개변수이름  매개변수형)반환형{} 입니다.
// func main(){}  =>  함수
// 반환값은 꼭 한개가 아닐수 도 있다.
// 반환 값은 retrun 값

// 매개변수가 있고, 반환 값도 있는 형태
// 매개변수가 있고, 반환 값이 없는 형태
// 매개변수가 없고, 반환 값이 있는 형태
// 매개변수가 없고, 반환 값이 없는 형태

package main

import "fmt"

func Hi(){//매개변수 X ,반환값 X.
	fmt.Println("앞에 사과가 있는데 당신은 사과를 드실겁니까?")
	fmt.Println("사과는 a박스에b개씩 담겨있다")
	fmt.Println("각각 몇박스 몇개인지 입력하시오:")
}

func input()(int,int){//매개변수x , 반환값 O
	var a, b int
	fmt.Scan(&a,&b)
	return a,b 
}

func result(c,d int)int{//매개변수O, 반환값O
	return c*d
}

func eatapple(num int){//매개변수 O, 반환값 X
	fmt.Printf("당신이 먹은 사과의 총 개수는 %d입니다.",num)
}

func main() {
	Hi()// 함수 실행
	apple1,apple2 := input()// 함수에 입력된 값을 담을 변수가 필요하다.
	realresult := result(apple1,apple2)
	eatapple(realresult)
}

// 1. Hi()함수는 매개변수도 없고 반환값도 없어 main에서 호출시 그냥 실행된다.
// 2. input()함수는 매개변수는 없고 반환값은 int형으로 선언되있다 scan으로 입력콘솔로서 작용하고 main함수에서 apple1,apple2변수를 통해 받는다.
// 3. result()함수는 매개변수,반환값 둘다 있고 main함수에서 apple1,apple2를 받아 c*d연산후 값을 반환하고 realresult변수에 대입
// 4. eatapple()함수는 매개변수는 있고 반환값은 없다. main함수에서 realresult를 받아 최종 문장을 출력한다.