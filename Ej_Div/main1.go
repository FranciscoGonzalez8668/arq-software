package main

import (
	"errors"
	"fmt"
)

func division(a float32, b float32) (float32, error) {
	var err error = nil
	if b == 0 {
		err = errors.New("Second value can not be 0")
		return -1, err
	}
	resultado := a / b
	return resultado, err

}

func main() {
	div, err := division(5, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(div)
	}
}
