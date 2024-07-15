package main

import (
	"fmt"
	"time"
)

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	//Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for i := range 3 {
		fmt.Println("range", i)
	}

	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	//You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// IF ELSE

	//basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//You can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//Logical operators like && and || are often useful in conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
	if num := 900; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//Note that you don’t need parentheses around conditions in Go, but that the braces are required.
	//There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

	// SWITCH
	// basic switch
	option := 1
	fmt.Print("Write ", option, " as ")
	switch option {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon", t.Month(), t.Day(), t.Hour())
	default:
		fmt.Println("It's after noon", t.Month(), t.Day(), t.Hour())
	}

	//A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause
	whatAmI := func(i interface{}) {
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
