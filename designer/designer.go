package designer

import "awesomeProject/employee"

type Designer struct {
	employee.Employee
}

func NewDesigner(position string, salary float64, address string) *Designer {
	return &Designer{employee.NewEmployee(position, salary, address)}
}
