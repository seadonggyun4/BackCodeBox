package main

import "fmt"

type dishOfSlice interface { //외부 공개 메서드 = 관계
	String() string //PutOfSlice은 String()이라는 외부 공개 메서드
}

type Slice interface { //오로지 관계만 선언해본다  ->jam을 slice로
	GetOnedish() dishOfSlice //Jam은 GetOnsSpoon()이라는 외부 공개 메서드 getonetons->getonedish
}

type sashimi struct { //Bread 객체를 만든다 -> 빅버거를 사시미로
	val string
}

func (b *sashimi) PutSlice(Slice Slice) { //오렌지잼이든 스토우베리 잼이든 상관없다. put patty ->putSlice로
	dish := Slice.GetOnedish() //tongs -> dish
	b.val += dish.String()
}

func (b *sashimi) String() string { //Bread 메서드2 String
	return "제주 모듬회 :  " + b.val
}

//======================================================================
type tunaSlice struct {
}

func (j *tunaSlice) GetOnedish() dishOfSlice {
	return &dishOftunaSlice{}
}

type FlatfishSlice struct {
}

func (j *FlatfishSlice) GetOnedish() dishOfSlice {
	return &dishOfFlatfishSlice{}
}

type salmonSlice struct {
}

func (j *salmonSlice) GetOnedish() dishOfSlice {
	return &dishOfsalmonSlice{}
}

type dishOftunaSlice struct {
}

func (s *dishOftunaSlice) String() string {
	return "+ 참치회"
}

type dishOfFlatfishSlice struct {
}

func (s *dishOfFlatfishSlice) String() string {
	return "+ 광어회"
}

type dishOfsalmonSlice struct {
}

func (s *dishOfsalmonSlice) String() string {
	return "+ 연어회"
}

func main() {
	sashimi := &sashimi{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	Slice := &salmonSlice{}
	Slice2 := &FlatfishSlice{}
	Slice3 := &tunaSlice{}
	sashimi.PutSlice(Slice)
	sashimi.PutSlice(Slice2)
	sashimi.PutSlice(Slice3)

	fmt.Println(sashimi)
}
