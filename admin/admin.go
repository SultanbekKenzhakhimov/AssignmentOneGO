package admin

import "awesomeProject/employee"

type Admin struct {
	employee.Employee
}

func NewAdmin(position string, salary float64, address string) *Admin {
	return &Admin{employee.NewEmployee(position, salary, address)}
}
