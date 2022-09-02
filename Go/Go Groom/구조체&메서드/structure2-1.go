package main

import "fmt"

type person struct{
	data	map[int]string
}

func newperson() *person{
	d := person{}
	d.data = map[int]string{}
	return &d
}


func main() {
	A1 := newperson()
	A1.data[1] = "10"
	A1.data[2] = "20"
	fmt.Println(A1)

}