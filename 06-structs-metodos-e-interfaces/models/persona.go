package models

import "fmt"

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      uint
	Direccion string
}

type PersonaConDireccion struct {
	Nombre    string
	Apellido  string
	Edad      uint
	Direccion Direccion
}

func NewPersonaConDireccion(nombre, apellido string, edad uint, direccion Direccion) *PersonaConDireccion {
	return &PersonaConDireccion{
		Nombre:    nombre,
		Apellido:  apellido,
		Edad:      edad,
		Direccion: direccion,
	}
}

type Asistente interface {
	AsisteAlCurso(curso string)
}

type PersonaConTelefono struct {
	Nombre   string
	Apellido string
	Edad     uint
	Telefono
}

func (pct *PersonaConTelefono) GetNombreCompleto(formato uint) string {
	if formato == 1 {
		return fmt.Sprintf("%s, %s", pct.Apellido, pct.Nombre)
	}
	if formato == 2 {
		return fmt.Sprintf("%c. %s", pct.Nombre[0], pct.Apellido)
	}
	return fmt.Sprintf("%s %s", pct.Nombre, pct.Apellido)
}

func (pct *PersonaConTelefono) AsisteAlCurso(curso string) {
	fmt.Printf("%s está asistiendo al curso %s\n", pct.Nombre, curso)
}
