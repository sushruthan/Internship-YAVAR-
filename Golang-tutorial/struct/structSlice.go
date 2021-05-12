package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}
type rectangle struct {
	length  int
	breadth int
	color   string

	geometry struct {
		area      int
		perimeter int
	}
}

func rectanglePrint() {
	var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}

func jasonStringPrint() {
	json_string := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

	emp1 := new(Employee)
	json.Unmarshal([]byte(json_string), emp1)
	fmt.Println(emp1)

	emp2 := new(Employee)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)
}
func slicePrint() {
	a := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

	b := make([]int, 5, 10) // Create a bigger slice
	copy(b, a)              // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}
func functionMake() { //declare a function using make for map
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)

	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)
}
func turnicateMap() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	// Method - I
	for k := range employee {
		delete(employee, k)
	}

	// Method - II
	employee = make(map[string]int)
}
func main() {
	jasonStringPrint()
	rectanglePrint()
	slicePrint()
	functionMake()
	turnicateMap()

}
