package main

import "fmt"

func main() {

	salary := map[string]int{"gabriel": 100, "george": 200, "junior": 300}
	fmt.Println(salary["gabriel"])
	// delete(salary, "gabriel")
	// salary["ga"] = 5000
	// fmt.Println(salary["ga"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["gabriel"] = 100

	for name, salary := range salary {
		fmt.Printf("the Salary de %s is %d \n", name, salary)
	}

	for _, salary := range salary {
		fmt.Printf("the salary %d \n", salary)
	}
}
