//[slice(슬라이스)]
//고정된 크기를 미리 지정하지 않고 이후에 필요에 따라 크기를 동적으로 변경가능
//슬라이스는  " var a []int " 와 같이 선언 -> 배열의 일부분을 가리키는 포인터를 만듭니다.
//슬라이스는 포인터, 실직절으로 어떤 변수가 들어갈 공간은 생성되지 않는다.
//배열은 복사할때 그 배열 내부 값들을 그대로 복사해 오지만, 슬라이스복사는 슬라이스가 참조하고있던 주소를 참조 하게 되는것이다.

//슬라이스는 아무런 값도  " var a []int " 와 같이 아무런 값도 지정안하면 0 이 아닌 크기도 용량도 없는 상태인 nil 값이 된다.

package main

import "fmt"

func main() {

	var a []int = []int{1,2,3}

	a[0] = 10 //슬라이스 인덱스 위치 선언후 값 초기화 가능

	fmt.Println(a)

	var b []int
	fmt.Println("b는 0이 아닌 nil 값 =>",b)


//======== Slice 는 make()함수로 선언만 하면서 크기를 미리 지정할 수 있습니다.=============
// " make(슬라이스 타입, 슬라이스 길이, 슬라이스의 용량) " 형태로 선언!!
// 용량은 생략하고 선언 가능하다! -> 길이와 동일하게 자동으로 용량 정해진다.
// make()함수를 사용해 슬라이스 메모리를 할당하고 난후 var을 이용해 값을 초기화 하면 새로운 메모리를 할당하면서 그전의 값은 없어집니다.
//메모리를 사용하고 값을 추가하기 위해서는 append()함수를 사용해야 합니다.
	
}