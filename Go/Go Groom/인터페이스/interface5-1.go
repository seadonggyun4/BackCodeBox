package main

import "fmt"

func main() {
	var empty interface{}  = 10
	a := empty//dinemic type
	b := empty.(int)//타입표명

	fmt.Println(a)
	printtest(b)

}

func printtest (i interface{}){
	fmt.Printf("%T,%d",i,i)
}