package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\nEnter the operation value (1 - 5) you want to perform 1. Add, 2. Subtract, 3. Mul, 4. Divide, 5. Quit:")
		var operator int
		_, err := fmt.Scanf("%d", &operator)
		if err != nil {
			fmt.Println(err)
			break
		}

		stringInput, err := inputParamPrinter(operator)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Print(stringInput)
			fmt.Println("\nEnter value 1:")
			var value1 float32
			_, err = fmt.Scanf("%f", &value1) // Use %f for float values
			if err != nil {
				fmt.Println("Please Provide the correct type of Integer or Float: ", err)
				break
			}

			fmt.Println("\nEnter value 2:")
			var value2 float32
			_, err = fmt.Scanf("%f", &value2) // Use %f for float values
			if err != nil {
				fmt.Println("Please Provide the correct type of Integer or Float: ", err)
				break
			}

			calculatorMain := Calculator{operator, value1, value2}

			result, err := operation(&calculatorMain)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("\nThe Result of the operation is =>", result, "\n========================================================================")

		}

	}
}
