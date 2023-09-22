package hr

import "awesomeProject/employee"

type Hr struct {
	employee.Employee
}

func NewHr(position string, salary float64, address string) *Hr {
	return &Hr{employee.NewEmployee(position, salary, address)}
}
