/* OOP 방식으로 빵에 잼 바르기 !!*/
package main

import "fmt"

type Bread struct { //bread객체를 만든다.
	val string
}

type StrawberryJam struct {
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string { //  SpoonOfStrawberryJam의 1개의 메소드
	return "+ strawberry"
}

func (j *StrawberryJam) GetOneSpoon() *SpoonOfStrawberryJam { // StrawberryJam의 1개 메서드 GetOneSpoon
	return &SpoonOfStrawberryJam{}
}

func (b *Bread) PutJam(jam *StrawberryJam) { //Bread 메서드1 string
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string { //Bread 메서드2 string
	return "bread" + b.val
}

func main() {
	bread := &Bread{}
	jam := &StrawberryJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
