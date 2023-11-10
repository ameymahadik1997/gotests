package main

import (
	"errors"
)

type Calculator struct {
	operator int
	value1   float32
	value2   float32
}

func operation(c *Calculator) (float32, error) {

	switch c.operator {
	case 1:
		return c.value1 + c.value2, nil
	case 2:
		return c.value1 - c.value2, nil
	case 3:
		return c.value1 * c.value2, nil
	case 4:
		return c.value1 / c.value2, nil
	default:
		return 0, errors.New("Invalid Option")
	}
}

func inputParamPrinter(operator int) (string, error) {
	switch operator {
	case 1:
		return "=============The operation to be performed is Addition=============", nil
	case 2:
		return "=============The operation to be performed is Substraction=============", nil
	case 3:
		return "=============The operation to be performed is Multiplication=============", nil
	case 4:
		return "=============The operation to be performed is Divide=============", nil
	case 5:
		return "", errors.New("Exiting The Calulator, Bye!")
	default:
		return "", errors.New("Invalid Option, Chose Again")
	}
}
