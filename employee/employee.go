package employee

type Employee interface {
	GetPosition() string
	SetPosition(string)
	GetSalary() float64
	SetSalary(float64)
	GetAddress() string
	SetAddress(string)
}

type MainEmployee struct {
	position string
	salary   float64
	address  string
}

func NewEmployee(position string, salary float64, address string) Employee {
	return &MainEmployee{position, salary, address}
}
func (m *MainEmployee) GetPosition() string {
	return m.position
}

func (m *MainEmployee) SetPosition(pos string) {
	m.position = pos
}

func (m *MainEmployee) GetSalary() float64 {
	return m.salary
}

func (m *MainEmployee) SetSalary(sal float64) {
	m.salary = sal
}

func (m *MainEmployee) GetAddress() string {
	return m.address
}

func (m *MainEmployee) SetAddress(addr string) {
	m.address = addr
}
