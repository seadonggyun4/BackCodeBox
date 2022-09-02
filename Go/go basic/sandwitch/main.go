/*절차지향적으로 빵을 만들기*/
package main

import "fmt"

type Bread struct {
	val string
}

//딸기잼
type StrawberryJam struct {
	opened bool // bool 타입은 참거짓을 나타내는 자료형
}

//오랜지 잼
type OrangeJam struct {
	opened bool // bool 타입은 참거짓을 나타내는 자료형
}

//SpoonOfStrawberry
type SpoonOfStrawberry struct {
}

//SpoonOfOrange
type SpoonOfOrange struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

//OpenStrawberryJam
func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

//OpenOrangeJam
func OpenOrangeJam(jam *OrangeJam) {
	jam.opened = true
}

// GetOneSpoon => 원하는 잼 타입에 따라 파라미터 수정 가능 : 딸기잼, 오랜지잼
func GetOneSpoon(_ *OrangeJam) *SpoonOfOrange {
	return &SpoonOfOrange{}
}

//PutJamOnBread => 원하는 잼 타입에 따라 파라미터 수정 가능 : 딸기잼, 오랜지잼
func PutJamOnBread(bread *Bread, jam *SpoonOfOrange) {
	bread.val += "+ Orange Jam "
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + "+"
	}
	return sandwitch
}

func main() {

	// 1. 빵 두개 꺼낸다.
	breads := GetBreads(2)

	// 원하는 잼 타입에 따라 수정
	jam := &OrangeJam{}

	// 2. 잼 뚜껑을 연다.
	OpenOrangeJam(jam)

	// 3. 잼 한스푼을 뜬다.
	spoon := GetOneSpoon(jam)

	// 4. 잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)

}
