package main

import "fmt"

//인터페이스는 메소드 묶음
type PutFoodOnBread interface { //떠온 재료들을 string함수로 빵 구조체에 집어넣는 메소드들의 집합
	String() string //String 함수 string타입
}
type PickFood interface { //재료들을 한스푼씩 떠오는 메소드 집합
	GetOnePatty() PutFoodOnBread //GetOnePatty 이라는 함수는 PutFoodOnBread와 같은 스트링 타입
}
type Bread struct { //bread는 구조체
	val string //val는 bread 구조체내부 객체 이다.
}

func (b *Bread) Put(food PickFood) { //b는 Bread의 주솟값 참조
	spoon := food.GetOnePatty()
	b.val += spoon.String() //put함수는 GetOnePatty으로 받아온 string타입의 데이터들을 val에 차곡차곡 저장
}

func (b *Bread) String() string { //b는 Bread의 주솟값 참조
	return "동대구 빅버거(빵 " + b.val + " +빵)" //val == Bread 인데 앞뒤에 "빵" 문자열 추가 // 내부 값이 뭔지 알수 없다.
}

//===================================================================재료 추가 색션!!!===================================================================================
type onion struct { //양파 구조체 생성
}

func (j *onion) GetOnePatty() PutFoodOnBread { // 양파 주소참조 -> GetOnePatty(한스푼 떠온다.)함수
	return &SpoonOfonion{} //onin 에 spoonOfoinon값인 "+양파" 를 반환한다.
}

type cheese struct {
}

func (j *cheese) GetOnePatty() PutFoodOnBread {
	return &SpoonOfcheese{}
}

type cabbage struct {
}

func (j *cabbage) GetOnePatty() PutFoodOnBread {
	return &SpoonOfcabbage{}
}

type tomato struct {
}

func (j *tomato) GetOnePatty() PutFoodOnBread {
	return &SpoonOftomato{}
}

type meat_patty struct {
}

func (j *meat_patty) GetOnePatty() PutFoodOnBread {
	return &SpoonOfmeat_patty{}
}

type krabby_patty struct {
}

func (j *krabby_patty) GetOnePatty() PutFoodOnBread {
	return &SpoonOfkrabby_patty{}
}

type SpoonOfonion struct { //양파 한스푼 구조체 생성
}

func (S *SpoonOfonion) String() string { // SpoonOfonion에 "+양파" 를 반환한다.
	return "+양파"

}

type SpoonOfcheese struct {
}

func (S *SpoonOfcheese) String() string {
	return "+치즈"

}

type SpoonOfcabbage struct {
}

func (S *SpoonOfcabbage) String() string {
	return "+양배추"

}

type SpoonOftomato struct {
}

func (S *SpoonOftomato) String() string {
	return "+토마토"

}

type SpoonOfmeat_patty struct {
}

func (S *SpoonOfmeat_patty) String() string {
	return "+고기패티"

}

type SpoonOfkrabby_patty struct {
}

func (S *SpoonOfkrabby_patty) String() string {
	return "+게살패티"

}

//============================================================================================================================================================================================================================
func main() {
	fmt.Println("[동대구역 5팀 빅버거]")
	fmt.Println("==============================")
	fmt.Println("[주문내역]")
	bread := &Bread{} // 모든 재료들이 빵 구조체에서 조합된다.
	Food := &cabbage{}
	Food1 := &cheese{}
	Food2 := &onion{}
	Food3 := &tomato{}
	Food4 := &meat_patty{}
	Food5 := &krabby_patty{}
	bread.Put(Food)
	bread.Put(Food1)
	bread.Put(Food2)
	bread.Put(Food3)
	bread.Put(Food4)

	fmt.Println("1번 손님: ", bread) //조합된 구조체 bread 출력
	fmt.Println("2번 손님: ", bread) //조합된 구조체 bread 출력
	bread.Put(Food5)
	fmt.Println("3번 손님: ", bread) //조합된 구조체 bread 출력
	fmt.Println("4번 손님: ", bread) //조합된 구조체 bread 출력
	fmt.Println("==============================")

}

/*
인터페이스는 외부공개 기능  집합(인터페이스는 행동의 집합쯤으로 이해)
구조체는 개체들의 집합
예제 햄버거 생성에서는 인터페이스 pickFood(음식들을 집어 오는 행동 메서드), PutFoodOnBread(음식들을 빵에 집어넣는, 연결해주는 메서드)
=> 두개의 행동을 인터페이스로 묶어버리면서 양배추,양파 등 구조체들이 새로 생겨날때마다 메서드도 그에 맞게 바꿔줘야 했는데
공통된 행동,기능을 하는  메서드를 묶어버리면서 메서드를 바꿀 필요가 없다?

*/
