package main
import "fmt"

func switchcase() {
	var day string
	fmt.Println("Enter the day: ")
fmt.Scan((&day))
	switch day {
	case "monday":
		fmt.Println("Today is Monday")

	case "tuesday":
		fmt.Println("Today is Tuesday")
	case "wednesday":
		fmt.Println("Today is Wednesday")
	case "thursday":
		fmt.Println("Today is Thursday")
	case "friday":
		fmt.Println("Today is Friday")
	case "saturday":
		fmt.Println("Today is Saturday")
	case "sunday":
		fmt.Println("Today is Sunday")
			
	default:
		fmt.Println("Invalid day")
}
}