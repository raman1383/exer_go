package main

import "fmt"

func main() {
	number1 := 12
	number2 := 20

	if number1 == number2 {
		fmt.Printf("Result: %d == %d", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}

	dayOfWeek := 3

	switch dayOfWeek {

	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}
}
