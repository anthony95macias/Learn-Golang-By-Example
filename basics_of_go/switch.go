package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	fmt.Print("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its the weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its the after noon")
	}

	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am a int")
		default:
			fmt.Printf("dont know type %T\n", t)
		}
	}
	WhatAmI(true)
	WhatAmI(1)
	WhatAmI("hey")
}
