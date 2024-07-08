package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Prinln("set:", a)
	fmt.Prinln("get:", a[4])

	fmt.Prinln("len:", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl: ", b)
}