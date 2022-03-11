package main

import (
	"fmt"
	"time"
)

const s string = "constant string"

func main() {
	//variables
	var firstname string = "test"
	fmt.Println(firstname)
	var age, age1 int = 10, 2
	fmt.Println(age + age1)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	var f string = "jacob"
	g := "nathan"
	fmt.Println(f + g)

	// constants, they can not be modified
	const s string = "constant"
	const t string = "string"
	fmt.Println(s + " " + t)
	const num int = 10
	const num1 int = 1000e10
	//another thing to know there is int8 int16 int32 int64
	//always used to increase memory size
	fmt.Println(int64(num1 / num))

	//for loop
	for { //infinite loop
		fmt.Println("loop")
		break
	}
	for g := 1; g <= 5; g++ {
		fmt.Println(g)
	}

	// if statements
	number := 8
	if number%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//switch statements
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	switch time.Now().Weekday() { //this gets the now time and weekday
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	whatAmI := func(i interface{}) { //interface can be anything
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
