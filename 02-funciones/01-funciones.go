package main

import (
	"fmt"
	"strings"
)

func sumar(n1, n2 int) int {
	return n1 + n2
}

func dividir(n1, n2 int) (int, string) {
	if n2 == 0 {
		return 0, "No puedes dividir por 0"
	}
	return n1 / n2, ""
}

func dividirConError(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("no puedes dividir por 0")
	}
	return n1 / n2, nil
}

func separarNombreCompleto(nombreCompleto string) (nombre, apellido string) {
	nombreCompletoSeparado := strings.Split(nombreCompleto, " ")
	nombre = nombreCompletoSeparado[0]
	apellido = nombreCompletoSeparado[1]
	return
}

func main() {

	fmt.Println(sumar(1, 2))
	fmt.Println(dividir(10, 2))

	resultado, err := dividir(10, 2)
	if err != "" {
		fmt.Println("ERROR ->", err)
	} else {
		fmt.Println("RESULTADO ->", resultado)
	}

	resultado, err = dividir(10, 0)
	if err != "" {
		fmt.Println("ERROR ->", err)
	} else {
		fmt.Println("RESULTADO ->", resultado)
	}

	// var err2 error

	resultado, err2 := dividirConError(10, 0)
	if err2 != nil {
		fmt.Println("ERROR ->", err2)
	} else {
		fmt.Println("RESULTADO ->", resultado)
	}

	// fmt.Println(10 / 0)

	nombre, apellido := separarNombreCompleto("Mike Kozinski")
	fmt.Println(nombre)
	fmt.Println(apellido)

	// EJERCICIO
	// generarEmailEmpresa("Mike Kozinski", "meta") -> m.kozinsky@meta.com

}
