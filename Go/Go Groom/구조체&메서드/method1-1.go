package main

import "fmt"

type squre struct{
	width int
	height int
}

func(s squre) squreArea()int{
	return s.width * s.height
}

func main() {
	squreA := new(squre)
	squreA.width = 20
	squreA.height = 20
	squreArea := squreA.squreArea()
	fmt.Println(squreArea)
}