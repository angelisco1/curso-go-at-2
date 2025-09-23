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

func generarEmailEmpresa(nombreCompleto, nombreEmpresa string) string {
	nombre, apellido := separarNombreCompleto(nombreCompleto)
	// fmt.Println(nombre[0])

	// email := strings.ToLower(string(nombre[0])) + "." + strings.ToLower(apellido) + "@" + strings.ToLower(nombreEmpresa) + ".com"
	// email := fmt.Sprintf("%s.%s@%s.com", strings.ToLower(string(nombre[0])), strings.ToLower(apellido), strings.ToLower(nombreEmpresa))

	// email := fmt.Sprintf("%c.%s@%s.com", strings.ToLower(nombre)[0], strings.ToLower(apellido), strings.ToLower(nombreEmpresa))
	email := strings.ToLower(fmt.Sprintf("%c.%s@%s.com", nombre[0], apellido, nombreEmpresa))
	// return m.kozinski@meta.com
	return email
}

func sumatorio(numeros ...int) int {
	// numeros es un slice de enteros
	suma := 0
	for i := 0; i < len(numeros); i++ {
		suma = suma + numeros[i]
	}
	return suma
}

func getTicketLoteria(sorteo string, numeros ...string) string {
	strNums := strings.Join(numeros, " - ")
	return fmt.Sprintf("Sorteo %s: %s", sorteo, strNums)
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
	// generarEmailEmpresa("Mike Kozinski", "meta") -> m.kozinski@meta.com

	emailMike := generarEmailEmpresa("Mike Kozinski", "Meta")
	fmt.Println(emailMike)

	fmt.Println(sumatorio(1, 2))
	fmt.Println(sumatorio(1, 2, 3, 4, 5))
	fmt.Println(sumatorio(1, 2, 3, 4, 5, 1, 2, 2, 2, 3, 4, 5, 5, 5, 2, 9, 9, 8, 3, 4, 2))

	fmt.Println(getTicketLoteria("Euromillón", "3", "19", "28", "34", "35"))
	fmt.Println(getTicketLoteria("Primitiva", "3", "19", "28", "34", "35", "47"))

}
