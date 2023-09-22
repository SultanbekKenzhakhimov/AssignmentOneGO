package main

import (
	"awesomeProject/admin"
	"awesomeProject/architecturer"
	"awesomeProject/designer"
	"awesomeProject/developer"
	"awesomeProject/employee"
	"awesomeProject/hr"
	"fmt"
)

func main() {

	firstEmployee := developer.NewDeveloper("Developer", 700000, "Zharokov st 179")
	showEmplDescription(firstEmployee)
	secondEmployee := designer.NewDesigner("Designer", 500000, "Zharokov st 179")
	showEmplDescription(secondEmployee)
	thirdEmployee := hr.NewHr("HR", 400000, "Zharokov st 179")
	showEmplDescription(thirdEmployee)
	fourthEmployee := admin.NewAdmin("Admin", 300000, "Zharokov st 179")
	showEmplDescription(fourthEmployee)
	fifthEmployee := architecturer.NewArchitecturer("Architecturer", 700000, "Zharokov st 179")
	showEmplDescription(fifthEmployee)
}
func showEmplDescription(e employee.Employee) {
	fmt.Printf("Position: %s\n", e.GetPosition())
	fmt.Printf("Salary: Tenge %.2f\n", e.GetSalary())
	fmt.Printf("Address: %s\n", e.GetAddress())
	fmt.Println()
}
