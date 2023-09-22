package architecturer

import "awesomeProject/employee"

type Architecturer struct {
	employee.Employee
}

func NewArchitecturer(position string, salary float64, address string) *Architecturer {
	return &Architecturer{employee.NewEmployee(position, salary, address)}
}
