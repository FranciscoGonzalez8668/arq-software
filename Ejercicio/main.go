package main

import (
	"fmt"
	"os"
)

func main() {
	nombre := "./texto.txt"
	file := crear(nombre)
	defer cerrar(file)
	escribir(file)

}

func crear(fileName string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}

func escribir(file *os.File) {

	fmt.Println("writing")
	fmt.Fprintln(file, "La revolucion francesa es un problema importante \nen cuanto a la programagiovanni de viviendas")

}

func cerrar(file *os.File) {
	fmt.Println("Closing")
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
