package main

import (
	"fmt"
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

}
