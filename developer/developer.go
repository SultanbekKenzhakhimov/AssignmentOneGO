package developer

import "awesomeProject/employee"

type Developer struct {
	employee.Employee
}

func NewDeveloper(position string, salary float64, address string) *Developer {
	return &Developer{employee.NewEmployee(position, salary, address)}
}
