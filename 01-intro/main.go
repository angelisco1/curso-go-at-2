package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println("Hola mundo!!!")

	// Variables
	var num int = 2
	fmt.Println(num)

	texto := "Hola mundo!!!"
	fmt.Println(texto)

	var esVerdad bool

	texto = "Adios"

	esVerdad = true
	fmt.Println(esVerdad)

	var x, y int = 1, 2
	var a, b string

	fmt.Println(x, y)

	a, b = "a", "b"
	fmt.Println(a, b)

	// Constantes
	const Token = "asdhj12ejk12jeh-12jk1h2ud-1d2jk-j12kdj"
	fmt.Println(Token)
	// Token = "asd"

	const (
		Lunes     = "L"
		Martes    = "M"
		Miercoles = "X"
		Jueves    = "J"
	)

	fmt.Println(Lunes)

	// Tipos de datos
	// - Numeros:
	// -- int, int8, int16, int32, int64
	// -- uint, uint8, uint16, uint32, uint64
	// -- float32, float64
	// -- complex64, complex128

	// int8 -> -127 a 128
	// uint8 -> 0 a 255 -> 2^8 = 256

	var rojo, verde, azul uint8 = 255, 10, 174
	fmt.Println(rojo, verde, azul)
	// azul = 490

	// var idUsuario uint64 = 120931748107412074
	var idUsuario uint = 120931748107412074
	fmt.Println(idUsuario)

	precio := 9.99
	fmt.Println(precio)

	cantidad := 3

	totalCompra := precio * float64(cantidad)
	fmt.Println(totalCompra)

	// Paquete math
	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Pow(2, 64))
	fmt.Println(math.Min(2, -64))

	var doce string = fmt.Sprintf("%d", 12)
	fmt.Println(doce)

	// - Strings -> con comillas dobles "" o con ``
	var nombre string = "Charly\n"
	fmt.Println(nombre)

	var apellido string = `Falco\n`
	fmt.Println(nombre, apellido)

	// - Rune es un int32 - representa un caracter Unicode
	var emoji rune = '🍺'
	fmt.Println(emoji)

	var emailRegexp string = `[a-zA-Z\d\.]+@[a-zA-Z\d\.]+\.[a-zA-Z]{2,}`
	fmt.Println(emailRegexp)
	emailRegexp = "[a-zA-Z\\d\\.]+@[a-zA-Z\\d\\.]+\\.[a-zA-Z]{2,}"
	fmt.Println(emailRegexp)

	var datos []byte = []byte("Hola mundo")
	fmt.Println(datos)

	fmt.Println(len(texto))

	// Paquete strings
	fmt.Println(strings.ToUpper(texto))
	fmt.Println(strings.ToTitle("nothing phone 3"))

	fmt.Println("  charly falco ")
	fmt.Println(strings.TrimSpace("  charly falco "))

	fmt.Println(strings.ReplaceAll("123,teclado,12.90€", "€", ""))

	filaCsv := strings.ReplaceAll("123,teclado,12.90€", ",", ";")
	fmt.Println(filaCsv)

	datosCsv := strings.Split(filaCsv, ";")
	fmt.Printf("El producto '%s' tiene el id %s y cuesta %s\n", datosCsv[1], datosCsv[0], datosCsv[2])

	// Los strings son inmutables
	// texto[1] = "z"

	fmt.Printf("Mi color favorito es: %X\n", 16431834)

	precio1 := 19.99
	precio2 := 579.9599
	precio3 := 1.2023

	fmt.Printf("%8.2f\n", precio1)
	fmt.Printf("%8.2f\n", precio2)
	fmt.Printf("%8.2f\n", precio3)

	// Booleanos
	var otroString string = ""
	if otroString != "" {
		fmt.Println("Tiene valor")
	} else {
		fmt.Println("No tiene valor")
	}

	// Valor nulo -> nil

	// Valores zero
	var numSinInicializar uint
	fmt.Println(numSinInicializar)

	var decimalSinInicializar float64
	fmt.Println(decimalSinInicializar)

	var textoSinInicializar string
	fmt.Println(textoSinInicializar)

	// Condicionales

	rolUsuario := "PREMIUM"
	// rolUsuario := "FREE"

	if rolUsuario == "ADMIN" {
		fmt.Println("Eres un administrador, puedes hacer de todo en la aplicación.")
	} else if rolUsuario == "PREMIUM" {
		fmt.Println("Eres un usuario de pago, puedes hacer operaciones de lectura y escritura en la aplicación.")
	} else if rolUsuario == "FREE" {
		fmt.Println("Eres un usuario que no paga, solo puedes ver los datos en la aplicación.")
	} else {
		fmt.Println("No tienes un rol asignado.")
	}

	switch rolUsuario {
	case "ADMIN":
		fmt.Println("Eres un administrador, puedes hacer de todo en la aplicación.")
	case "PREMIUM":
		fmt.Println("Eres un usuario de pago, puedes hacer operaciones de lectura y escritura en la aplicación.")
	case "FREE", "FREEMIUM":
		fmt.Println("Eres un usuario que no paga, solo puedes ver los datos en la aplicación.")
	default:
		fmt.Println("No tienes un rol asignado.")
	}

	// Bucles

	for i := 0; i < 6; i++ {
		if i == 3 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	j := 0
	for j < 4 {
		j++
	}

}
