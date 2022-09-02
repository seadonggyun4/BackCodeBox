package main

import "fmt"

type SpoonOfJam interface { //외부 메서드 = 관계
	String() string //SpoonOfJam 은 string()이라는 외부 메서드
}
type Jam interface { //오로지 관계만 선언해 본다.
	GetOneSpoon() SpoonOfJam //Jam은 GetoneSpoon()이라는 외부 메서드
}
type Bread struct { //bread객체 생성
	val string
}

func (b *Bread) PutJam(jam Jam) { // 여러타입의 잼 생성가능
	spoon := jam.GetOneSpoon() //spoon객체에 jam 함수의 GetOneSpoon를 할당
	b.val += spoon.String()    //b.val 에 spoon값을 추가
}

func (b *Bread) String() string { // 추가 bread 값과 putjame함수로 추가된 값을 을 struct에 추가
	return "bread" + b.val
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct {
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfStrawberryJam struct {
}

func (S *SpoonOfStrawberryJam) String() string {
	return "+strawberry"

}

type SpoonOfOrangeJam struct {
}

func (S *SpoonOfOrangeJam) String() string {
	return "+Orange"

}

type SpoonOfAppleJam struct {
}

func (S *SpoonOfAppleJam) String() string {
	return "+Apple"

}

func main() {
	bread := &Bread{}
	jam := &AppleJam{}
	jam1 := &OrangeJam{}
	jam2 := &StrawberryJam{}
	bread.PutJam(jam)
	bread.PutJam(jam1)
	bread.PutJam(jam2)

	fmt.Println(bread)

}
