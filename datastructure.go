package main

import "fmt"

type Employee struct{
	id   int
	name string
	department string	
}

func main(){
	var employee = new(Employee)
	employee.id = 100
	employee.name = "Balaji"
	employee.department = "IT"
	
	fmt.Println(employee)	
}