package main

import (
	"fmt"
	"os"
)

func main() {

	// For usando range que devuelve dos parametros pos, char
	for pos, char := range "日本\x80語" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	//Defer sirve para ejecutar las cosas al final de la funcion(invierte el orden es como una pila)
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
	//File write
	f, _ := os.Create("D:/Facultad/3ro/Arquitectura de sofware/ejercicios/Ej_for/x.txt")
	fmt.Fprintln(f, "data")
	f.Close()

	//file reader
	file2, erro := os.ReadFile("D:/Facultad/3ro/Arquitectura de sofware/ejercicios/Ej_for/x.txt")
	fmt.Print(string(file2))
	if erro != nil {
		return
	}
	//erro = file2.Close()
}
