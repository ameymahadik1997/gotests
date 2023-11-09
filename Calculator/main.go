package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter the operation value (1 - 5) you want to perform 1. Add, 2. Subtract, 3. Mul, 4. Divide, 5. Quit:")
	var operator int
	_, err := fmt.Scanf("%d", &operator)
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	fmt.Println("Enter value 1:")
	var value1 float32
	_, err = fmt.Scanf("%f", &value1) // Use %f for float values
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	fmt.Println("Enter value 2:")
	var value2 float32
	_, err = fmt.Scanf("%f", &value2) // Use %f for float values
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	calculatorMain := Calculator{operator, value1, value2}

	result, err := operation(&calculatorMain)
	if err != nil {
		log.Fatalf("Calculation error: %v", err)
	}

	fmt.Println("The Result of the operation is =>", result)
}
