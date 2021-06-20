package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return nil
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Robson",
			ID:   "123456",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())
}
