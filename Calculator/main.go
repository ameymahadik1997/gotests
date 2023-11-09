package main

import (
	"fmt"
)

func main() {

	// fmt.Println("Enter the operation value(1 - 5) you want to perform 1. Add, 2. Substract, 3. Mul, 4. Divide, 5. Quit:")
	// var operator int
	// _, err := fmt.Scanf("%d", &operator)
	// if err != nil {
	// 	fmt.Errorf("Something went Wrong")
	// }

	fmt.Println("Enter value 1:")
	var value1 float32
	_, err := fmt.Scanf("%d", &value1)
	if err != nil {
		fmt.Errorf("Something went Wrong")
	}

	fmt.Println("Enter value 2:")
	var value2 float32
	_, err = fmt.Scanf("%d", &value2)
	if err != nil {
		fmt.Errorf("Something went Wrong")
	}

	calculatorMain := Calculator{3, 2, 2}

	result, err := operation(&calculatorMain)
	if err != nil {
		fmt.Errorf("Error")
	}

	fmt.Println("The Result of the operation is => ", result)
}
