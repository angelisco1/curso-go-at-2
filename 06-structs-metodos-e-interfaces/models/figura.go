package models

import "math"

type Figura interface {
	CalcularArea() float64
}

type Rectangulo struct {
	Lado1 float64
	Lado2 float64
}

func (r *Rectangulo) CalcularArea() float64 {
	return r.Lado1 * r.Lado2
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) CalcularArea() float64 {
	return math.Pi * math.Pow(c.Radio, 2.0)
}
