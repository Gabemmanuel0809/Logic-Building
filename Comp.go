package main 
import "fmt"

type Employee struct {
	Id int 
	Name string 
	Role string 
}

type Company struct {
	Employees []Employee
}

func (c *Company) hire() {
	var id int 
	fmt.Println("Input the id number")
  fmt.Scanf("%d", &id)

	var name string 
	fmt.Println("Input the name")
	fmt.Scanf("%s", &name)

	var role string 
	fmt.Println("Input the role")
	fmt.Scanf("%s", &role)

	var newEmp = Employee {
		Id: id,
		Name: name,
		Role: role,
	}

  c.Employees = append(c.Employees, newEmp)
	fmt.Println("Hiring Success")

}

func main() {
   var comp = Company{
	    Employees: []Employee{},
   }

   comp.hire()

   fmt.Println(comp)
}
