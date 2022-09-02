package main

import "fmt"

//인터페이스는 메소드 묶음
//===================================================================위치 선정색션!!!===================================================================================

type ChoosePlaceInBusan interface { //GetOneSpoon과 같은 역할
	String() string
}
type PickPlace interface {
	getoneplace() ChoosePlaceInBusan
}

func (b *Busan) Choose(place PickPlace) { //put jam과 같은 역할
	poon := place.getoneplace()
	b.val += poon.String()
}

type Busan struct { //bread와 같은역할
	val string
}

func (b *Busan) String() string { // 값들을 조합하는 기능
	return "부산 " + b.val + " 모듬회 : "
}

//===================================================================회 선정 색션!!!===================================================================================
type PutSliceOnDish interface { //GetOneSpoon과 같은 역할
	String() string
}
type PickSlice interface {
	GetOneslice() PutSliceOnDish
}
type Dish struct { //bread와 같은역할
	val string
}

func (b *Dish) Put(food PickSlice) { //put jam과 같은 역할
	Slice := food.GetOneslice()
	b.val += Slice.String()
}

func (b *Dish) String() string { // 값들을 조합하는 기능
	return "(접시 " + b.val + " + 와사비)"
}

//===================================================================위치객체들 색션!!!===================================================================================
type Gwang_anli struct { //Jam=Patty와 같은 역할
}

func (j *Gwang_anli) getoneplace() ChoosePlaceInBusan {
	return &PlaceOfGwang_anli{}
}

type PlaceOfGwang_anli struct { //Spoon=tongs 과 같은 역할
}

func (S *PlaceOfGwang_anli) String() string {
	return " 광안리"

}

type Haeundae struct { //Jam=Patty와 같은 역할
}

func (j *Haeundae) getoneplace() ChoosePlaceInBusan {
	return &PlaceOfHaeundae{}
}

type PlaceOfHaeundae struct { //Spoon=tongs 과 같은 역할
}

func (S *PlaceOfHaeundae) String() string {
	return " 해운대"

}

//===================================================================회객체들 색션!!!===================================================================================

type tuna struct { //Jam=Patty와 같은 역할
}

func (j *tuna) GetOneslice() PutSliceOnDish {
	return &SliceOftuna{}
}

type SliceOftuna struct { //Spoon=tongs 과 같은 역할
}

func (S *SliceOftuna) String() string {
	return "+ 참치회"

}

type salmon struct { //Jam=Patty와 같은 역할
}

func (j *salmon) GetOneslice() PutSliceOnDish {
	return &SliceOfsalmon{}
}

type SliceOfsalmon struct { //Spoon=tongs 과 같은 역할
}

func (S *SliceOfsalmon) String() string {
	return "+ 연어회"

}

//============================================================================================================================================================================================================================
func main() {
	fmt.Println("[부산지역 모듬회]")
	fmt.Println("=====================================================================")
	fmt.Println("[주문내역]")
	Busan := &Busan{}
	Dish := &Dish{}

	place1 := &Gwang_anli{}
	place2 := &Haeundae{}
	sashimi1 := &tuna{}
	sashimi2 := &salmon{}

	Busan.Choose(place1)
	Dish.Put(sashimi1)
	fmt.Print("\n")
	fmt.Println("1번 손님: ", Busan, Dish)
	Busan.Choose(place2)
	fmt.Println("2번 손님: ", Busan, Dish)
	Dish.Put(sashimi2)
	fmt.Println("3번 손님: ", Busan, Dish)
	fmt.Println("4번 손님: ", Busan, Dish)
	fmt.Print("\n")
	fmt.Println("======================================================================")

}
