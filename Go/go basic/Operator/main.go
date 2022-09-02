package main

import "fmt"

/*
   소스, 패티/야채
*/

type StackOfIngredients interface {
   AddBread() string
}

type Ingredients interface {
   GetIngredients() StackOfIngredients
}

type SprinkleOfSauce interface {
   AddBread() string
}

type Sauce interface {
   GetSauce() SprinkleOfSauce
}

type 

// Bread 객체
type Bread struct { //bread객체 생성
	val string
}

func (b *Bread) AddIngredient(ing Ingredients) {
   ingredient := ing.GetIngredients()
   b.val += ingredient.AddBread()
}

// 빵
func (b *Bread) AddBread() string {
	return  "bread " + b.val 
}

type Patty struct{

}

type VolcanoSauce struct {
}

func (s *VolcanoSauce) GetSauce() SprinkleOfSauce {
   return &SprinkleOfVolcanoSauce{}
}

type SoyGarlicSauce struct {
}

func (s *SoyGarlicSauce) GetSauce() SprinkleOfSauce {
   return &SprinkleOfSoyGarlicSauce{}
}

type TomatoSauce struct {
}

func (s *TomatoSauce) GetSauce() SprinkleOfSauce {
   return &SprinkleOfTomatoSauce{}
}

type SprinkleOfVolcanoSauce struct {
}

func (s *SprinkleOfVolcanoSauce) AddBread() string {
   return "+ VolcanoSauce"
}

type SprinkleOfTomatoSauce struct {
}

func (s *SprinkleOfTomatoSauce) AddBread() string {
   return "+ TomatoSauce"
}

type SprinkleOfSoyGarlicSauce struct {
}

func (s *SprinkleOfSoyGarlicSauce) AddBread() string {
   return "+ SoyGarlicSauce"
}

func main() {
   fmt.Println("5팀 버거에 오신걸 환영합니다.")
   bread := &Bread{}
   Food := &TomatoSauce{}
   bread. AddIngredient(food)

   fmt.Println(bread)

//    Ingredients := &MeetPatty{}
}

// 음료
// type Drink struct {
//    val string
// }

// 사이드 메뉴
// type Side struct {
//    val string
// }

// // 세트 메뉴
// type SetMenu struct {
//    bread := &Bread{}
// }