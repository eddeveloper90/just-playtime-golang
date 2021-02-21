package main

import (
	"fmt"
	"time"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	fmt.Println(time.Now().Unix())
	i := 2
	fmt.Println("Write ", i, " as ")
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
			fmt.Println("It`s the weekend")
	default:
		fmt.Println("It`s a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon")
	}

	whatAmI := func(i interface{}) {
		describe(i)

		switch i.(type) {
		case bool:
			fmt.Println("I`m a bool.")
		case int:
			fmt.Println("I`m a int.")
		default:
			fmt.Println("Dont know the type")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}