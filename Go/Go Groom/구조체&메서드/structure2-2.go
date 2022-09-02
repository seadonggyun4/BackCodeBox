package main

import "fmt"

type animals struct{
	data map[int]string
}

func animalsData() *animals{
	d := animals{}
	d.data = map[int]string{}
	return &d
}

func main() {
	animaltype := animalsData()
	animaltype.data[1] = "dog"
	animaltype.data[2] = "cat"
	animaltype.data[10] = "cow"

	fmt.Println(animaltype)

}