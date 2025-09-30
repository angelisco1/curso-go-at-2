package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type MiFunc func(a, b int) int

func realizarOperacion(a, b int, fn MiFunc) int {
	return fn(a, b)
}

type FiltroFunc func(n int) bool

func FiltrarNumeros(nums []int, filtro FiltroFunc) []int {
	numerosFiltrados := []int{}
	for _, num := range nums {
		if filtro(num) {
			numerosFiltrados = append(numerosFiltrados, num)
		}
	}
	return numerosFiltrados
}

type OperacionFunc func(n int) int

func Map(nums []int, operacion OperacionFunc) []int {
	numerosTransformados := []int{}
	for _, num := range nums {
		nuevoNum := operacion(num)
		numerosTransformados = append(numerosTransformados, nuevoNum)
	}
	return numerosTransformados
}

type OperacionFunc2 func(n int) string

func MapIntStr(nums []int, operacion OperacionFunc2) []string {

	numerosTransformados := []string{}
	for _, num := range nums {
		nuevoNum := operacion(num)
		numerosTransformados = append(numerosTransformados, nuevoNum)
	}
	return numerosTransformados
}

// Reflection API - filter generico

func Filter(lista interface{}, fn interface{}) interface{} {

	tipoLista := reflect.TypeOf(lista)
	tipoFn := reflect.TypeOf(fn)

	fmt.Println(tipoLista.Kind())
	fmt.Println(tipoFn.Kind())

	if tipoLista.Kind() != reflect.Slice {
		panic("debería de ser un slice")
	}

	if tipoFn.Kind() != reflect.Func {
		panic("debería de ser una función")
	}

	valorLista := reflect.ValueOf(lista)
	valorFn := reflect.ValueOf(fn)

	fmt.Println("->", valorLista)

	listaFiltrados := reflect.MakeSlice(tipoLista, 0, valorLista.Len())

	for i := 0; i < valorLista.Len(); i++ {
		item := valorLista.Index(i)
		resultadoSlice := valorFn.Call([]reflect.Value{item})
		resultado := resultadoSlice[0].Bool()
		if resultado {
			listaFiltrados = reflect.Append(listaFiltrados, item)
		}
	}

	return listaFiltrados
}

type Persona struct {
	Nombre string
	Edad   uint
}

func FilterConGenericos[T any](lista []T, fn func(T) bool) []T {
	elementosFiltrados := []T{}

	for _, item := range lista {
		if fn(item) {
			elementosFiltrados = append(elementosFiltrados, item)
		}
	}

	return elementosFiltrados
}

// Closures

func Contador() func() int {

	cuenta := 0

	return func() int {
		cuenta++
		return cuenta
	}
}

func accederPorRol(rol string) func(recurso string) bool {
	return func(recurso string) bool {
		if recurso == "/admin" && rol == "ADMIN" {
			return true
		}

		if recurso == "/facturas" && (rol == "PREMIUM" || rol == "ADMIN") {
			return true
		}

		if recurso == "/home" {
			return true
		}

		return false
	}
}

func crearValidador(longitudCampo int) func(valorCampo string) bool {
	return func(valorCampo string) bool {
		return len(valorCampo) >= longitudCampo
	}
}

func main() {

	dobleNums := Map([]int{1, 2, 3, 4, 5}, func(n int) int {
		return n * 2
	})
	fmt.Println(dobleNums)

	numsStr := MapIntStr([]int{1, 2, 3, 4, 5}, func(n int) string {
		// return fmt.Sprintf("%d!", n)
		return strconv.Itoa(n) + "!"
	})
	fmt.Println(numsStr)

	numsCuyoDobleEsMayorQue6 := FiltrarNumeros([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n*2 > 6
	})
	fmt.Println(numsCuyoDobleEsMayorQue6)

	var filtroNumsPares FiltroFunc = func(n int) bool {
		return n%2 == 0
	}

	numsPares := FiltrarNumeros([]int{1, 2, 3, 4, 5}, filtroNumsPares)
	fmt.Println(numsPares)

	var sumar MiFunc = func(a, b int) int {
		return a + b
	}

	restar := func(a, b int) int {
		return a - b
	}

	fmt.Println(sumar(1, 2))
	fmt.Println(sumar(2, 2))

	fmt.Println(realizarOperacion(1, 2, sumar))
	fmt.Println(realizarOperacion(1, 2, restar))

	contador1 := Contador()

	fmt.Println(contador1())
	fmt.Println(contador1())
	fmt.Println(contador1())
	fmt.Println(contador1())

	accesoAdmin := accederPorRol("ADMIN")
	accesoFree := accederPorRol("FREE")

	fmt.Println(accesoAdmin("/home"))
	fmt.Println(accesoAdmin("/admin"))

	fmt.Println(accesoFree("/home"))
	fmt.Println(accesoFree("/admin"))

	validarContrasenya := crearValidador(8)
	validarNombre := crearValidador(3)

	fmt.Println(validarContrasenya("qwerty"))
	fmt.Println(validarContrasenya("qwerty12345"))
	fmt.Println(validarNombre("Aa"))
	fmt.Println(validarNombre("Aaaa"))

	// Reflection API

	var mensaje interface{}
	mensaje = "Un string"

	tipo := reflect.TypeOf(mensaje)
	valor := reflect.ValueOf(mensaje)
	fmt.Println("----")
	fmt.Println(tipo)
	fmt.Println(valor)

	unNumero := 2
	tipoUnNumero := reflect.TypeOf(unNumero)
	fmt.Println(tipoUnNumero)
	valorUnNumero := reflect.ValueOf(unNumero)
	fmt.Println(valorUnNumero)

	r := Filter([]string{"a", "b", "a", "c"}, func(item string) bool {
		return item == "a"
	})
	fmt.Println(r)

	r = Filter([]int{1, 2, 3, 4}, func(item int) bool {
		return item > 2
	})
	fmt.Println(r)

	r = FilterConGenericos([]int{1, 2, 3, 4}, func(item int) bool {
		return item > 2
	})
	fmt.Println(r)

	r = FilterConGenericos([]Persona{
		Persona{"Charly", 48},
		Persona{"Mike", 53},
		Persona{"Sarah", 50},
		Persona{"Faustina", 46},
	}, func(item Persona) bool {
		return item.Edad >= 50
	})
	fmt.Println(r)

}
