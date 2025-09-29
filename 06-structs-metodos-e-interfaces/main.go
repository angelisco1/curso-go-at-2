package main

import (
	"fmt"
	"reflect"
	"structs/models"
	// . "structs/models"
	// _ "structs/models"
	// mod "structs/models"
)

func main() {

	charly := models.Persona{
		Nombre:    "Charly",
		Apellido:  "Falco",
		Edad:      48,
		Direccion: "C/ Norte 23, New York",
	}

	fmt.Println(charly)

	direccionDeMike := models.Direccion{
		Tipo:   "Calle",
		Nombre: "Este",
		Numero: 10,
		Ciudad: "New York",
	}

	mike := models.PersonaConDireccion{
		Nombre:    "Mike",
		Apellido:  "Kozinski",
		Edad:      53,
		Direccion: direccionDeMike,
	}

	fmt.Println(mike.Nombre)
	fmt.Println(mike.Apellido)
	fmt.Println(mike.Direccion.Nombre)

	direccion2 := models.Direccion{"Avenida", "La mancha", 152, "Toledo"}
	fmt.Println(direccion2)

	direccion3 := models.Direccion{
		// Tipo:   "Avenida",
		Nombre: "La mancha",
		Numero: 124,
	}
	fmt.Println(direccion3)

	direccionVacia := models.Direccion{}
	fmt.Println(direccionVacia)

	mike2 := models.PersonaConTelefono{
		Nombre:   "Mike",
		Apellido: "Kozinski",
		Edad:     53,
		Telefono: models.Telefono{
			Prefijo: "+34",
			Numero:  666777888,
		},
	}

	fmt.Println(mike2.Numero)

	fmt.Println(mike2.GetNombreCompleto(3))
	fmt.Println(mike2.GetNombreCompleto(1))
	fmt.Println(mike2.GetNombreCompleto(2))

	mike2.AsisteAlCurso("Golang")

	figuras := []models.Figura{
		&models.Rectangulo{
			Lado1: 2,
			Lado2: 3,
		},
		&models.Rectangulo{
			Lado1: 4,
			Lado2: 3,
		},
		&models.Circulo{
			Radio: 5,
		},
	}

	for _, figura := range figuras {
		fmt.Printf("El área de la figura %+w es: %.2f\n", figura, figura.CalcularArea())
	}

	// Interface vacía
	// var cualquierDato interface{}
	var cualquierDato any

	cualquierDato = 3
	fmt.Println(cualquierDato)

	cualquierDato = "Hola mundo"
	fmt.Println(cualquierDato)

	cualquierDato = figuras[0]
	cualquierDato = figuras[2]

	valorCualquierDato, ok := cualquierDato.(*models.Rectangulo)
	if ok {
		fmt.Println("Rectangulo: ", valorCualquierDato.Lado1, valorCualquierDato.Lado2)
	} else {
		fmt.Println("Esto no es un rectangulo")
		valorCualquierDato2, ok := cualquierDato.(*models.Circulo)
		if ok {
			fmt.Println("Circulo: ", valorCualquierDato2.Radio)
		} else {
			fmt.Println("Esto no es un circulo")
		}
	}

	var cualquierDato2 any
	cualquierDato2 = figuras[0]
	// cualquierDato2 = figuras[2]
	// cualquierDato2 = "Hola mundo"

	// fmt.Println(cualquierDato2.(type))

	switch valor := cualquierDato2.(type) {
	case *models.Circulo:
		fmt.Println("Es un circulo: ", valor)
	case *models.Rectangulo:
		fmt.Println("Es un rectangulo: ", valor)
	case string:
		fmt.Println("Es un string: ", valor)
	default:
		fmt.Println("No se lo que es: ", valor)
	}

	// Reflection API
	tipoMike := reflect.TypeOf(mike2)
	valorMike := reflect.ValueOf(mike2)
	fmt.Println(tipoMike)
	fmt.Println(valorMike)

	// fmt.Println(tipoMike.NumMethod())
	// fmt.Println(tipoMike.Method(0).Name)

	rect := &models.Rectangulo{
		Lado1: 4,
		Lado2: 3,
	}

	tipoRect := reflect.TypeOf(rect)
	fmt.Println(tipoRect.NumMethod())
	fmt.Println(tipoRect.Method(0).Name)

	valorRect := reflect.ValueOf(rect)
	metodoCalcularArea := valorRect.MethodByName("CalcularArea")
	area := metodoCalcularArea.Call([]reflect.Value{})
	fmt.Println(area[0])

}
