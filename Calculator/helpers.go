package main

import (
	"errors"
	"fmt"
)

type Calculator struct {
	operator int
	value1   float32
	value2   float32
}

func operation(c *Calculator) (float32, error) {
	fmt.Println(c.operator, c.value1, c.value2)
	switch c.operator {
	case 1:
		return c.value1 + c.value2, nil
	case 2:
		return c.value1 - c.value2, nil
	case 3:
		return c.value1 * c.value2, nil
	case 4:
		return c.value1 / c.value2, nil
	case 5:
		return 0, errors.New("Exiting")
	default:
		return 0, errors.New("Invalid Option")
	}
}
