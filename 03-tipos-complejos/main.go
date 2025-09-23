package main

import (
	"fmt"
	"maps"
	"slices"
)

func Shift(datos []int) ([]int, error) {
	if len(datos) == 0 {
		return nil, fmt.Errorf("no se puede aplicar esta operación a un slice vacío")
	}
	return datos[1:], nil
}

func Unshift(datos []int, elem int) []int {
	return append([]int{elem}, datos...)
}

func main() {
	// Arrays: tienen un tamaño fijo

	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println(nums)

	nums2 := [4]int{1, 2, 3}
	nums2[3] = 10
	fmt.Println(nums2)

	aprobados := [5]bool{}
	fmt.Println(aprobados)
	aprobados[2] = true
	aprobados[4] = true
	// aprobados[5] = true
	fmt.Println(aprobados)

	sugus := [...]string{"piña", "naranja", "fresa", "limón"}
	fmt.Println(sugus)

	// Slices: tienen un tamaño variable
	var letras []string = []string{"a", "b", "c"}
	fmt.Println(letras)
	fmt.Println(cap(letras))
	fmt.Println(len(letras))
	// letras[3] = "d"
	// fmt.Println(letras)

	bools := make([]bool, 3, 6)
	fmt.Println(bools)
	bools = append(bools, true, false, true)
	fmt.Println(bools)
	bools = append(bools, true)
	fmt.Println(bools)
	fmt.Println(len(bools))
	fmt.Println(cap(bools))

	bools = slices.Delete(bools, 3, 6)
	fmt.Println(bools)
	fmt.Println(len(bools))
	fmt.Println(cap(bools))

	letras2 := []string{"z", "f", "u"}

	letras3 := append(letras, letras2...)
	fmt.Println(letras3)

	copy(letras, letras2)
	fmt.Println(letras)

	for pos, letra := range letras3 {
		fmt.Printf("%d: %s\n", pos, letra)
	}

	// slicing
	fmt.Println(letras3[1:3])
	fmt.Println(letras3[:3])
	fmt.Println(letras3[3:])

	emailsHackeados := []string{
		"angel@gmail.com",
		"mike@gmail.com",
		"charly@gmail.com",
		"sarah@gmail.com",
		"octavia@gmail.com",
	}

	emailsHackeadosDeInternet := []string{
		"charly@gmail.com",
		"facundo@gmail.com",
		"maria@gmail.com",
		"sarah@gmail.com",
	}

	fmt.Println(slices.Contains(emailsHackeados, "facundo@gmail.com"))
	fmt.Println(slices.Contains(emailsHackeados, "mike@gmail.com"))

	emailsOrdenados := slices.Sorted(slices.Values(emailsHackeados))
	fmt.Println(emailsOrdenados)

	slices.Sort(emailsHackeados)
	fmt.Println(emailsHackeados)

	todosLosEmailsHackeados := append(emailsHackeados, emailsHackeadosDeInternet...)
	slices.Sort(todosLosEmailsHackeados)
	emailsHackeadosSinRepeticiones := slices.Compact(todosLosEmailsHackeados)
	fmt.Println(emailsHackeadosSinRepeticiones)

	// Maps

	configuracion := map[string]string{
		"puerto":   "8080",
		"password": "qwerty",
		"host":     "localhost",
	}
	fmt.Println(configuracion)
	fmt.Printf("%s:%s\n", configuracion["host"], configuracion["puerto"])

	configuracion["puerto"] = "3000"
	fmt.Printf("%s:%s\n", configuracion["host"], configuracion["puerto"])

	bbdd, existeClave := configuracion["bbdd"]
	if existeClave {
		fmt.Println("Existe:", bbdd)
	} else {
		fmt.Println("No existe la clave en el mapa")
	}

	carritoCompra := map[string]int{
		"naranjas":       2,
		"filete ternera": 3,
		"botella agua":   2,
	}

	for clave, valor := range carritoCompra {
		fmt.Printf("%s x %d\n", clave, valor)
	}

	for valor := range maps.Values(carritoCompra) {
		fmt.Println(valor)
	}

	// EJERCICIO: crear las funciones shift y unshift
	// shift([1, 2, 3, 4]) -> retorna [2, 3, 4]
	// unshift([2, 3, 4], 1) -> retorna [1, 2, 3, 4]
	fmt.Println(Shift([]int{1, 2, 3, 4}))
	fmt.Println(Unshift([]int{1, 2, 3, 4}, 0))

	subSlice, err := Shift([]int{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(subSlice)
	}
}
