package main

import "fmt"

// const (
// 	Lunes     int = 0
// 	Martes    int = 1
// 	Miercoles int = 2
// 	Jueves    int = 3
// 	Viernes   int = 4
// 	Sabado    int = 5
// 	Domingo   int = 6
// )

const (
	Lunes int = iota
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)

const (
	Amarillo int = iota + 1
	Verde
	Azul
	Naranja
)

const (
	_ int = iota
	Enero
	Febrero
	Marzo
)

type Error int

const (
	BadRequest   Error = 400
	Unauthorized       = 401
	NotFound           = 404
)

func (e Error) String() string {
	switch e {
	case 400:
		return "Bad request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not found"
	default:
		return "Internal server error"
	}
}

func getError(codigo int) string {
	switch codigo {
	case 400:
		return "Bad request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not found"
	default:
		return "Internal server error"
	}
}

func main() {
	fmt.Println(Miercoles)
	fmt.Println(Amarillo)
	fmt.Println(Azul)

	fmt.Println(BadRequest)
}
