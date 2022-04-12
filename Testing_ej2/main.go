package main

import "errors"

func Division(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("no puedo dividir por 0")
	}
	return a / b, nil
}
