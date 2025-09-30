package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func realizarTareas(tareas []string, canal chan string, wc *sync.WaitGroup) {
	defer wc.Done()

	for _, tarea := range tareas {
		canal <- tarea
	}
	close(canal)
}

func recibirTareasHechas(canal chan string, wc *sync.WaitGroup) {
	defer wc.Done()

	valor, ok := <-canal
	if !ok {
		fmt.Println("El canal está cerrado.")
	} else {
		fmt.Println("El canal está abierto: ", valor)
	}

	for tarea := range canal {
		fmt.Printf("%s: realizada\n", tarea)
	}

	valor, ok = <-canal
	if !ok {
		fmt.Println("El canal está cerrado.")
	} else {
		fmt.Println("El canal está abierto: ", valor)
	}
}

func getDatosApi1(canal chan string) {
	defer close(canal)
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	canal <- "Respuesta de la API 1"

}

func getDatosApi2(canal chan string) {
	defer close(canal)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	canal <- "Respuesta de la API 2"
}

func main() {
	// canalNumeros := make(chan int)

	// go func() {
	// 	canalNumeros <- 2

	// 	time.Sleep(3 * time.Second)
	// 	canalNumeros <- 18
	// 	canalNumeros <- 36
	// }()

	// num1 := <-canalNumeros
	// fmt.Println(num1)

	// num2 := <-canalNumeros
	// fmt.Println(num2)

	// num3 := <-canalNumeros
	// fmt.Println(num3)

	canalConBuffer := make(chan int, 2)

	go func() {
		canalConBuffer <- 1
		canalConBuffer <- 2
		canalConBuffer <- 3
		// time.Sleep(4 * time.Second)
		canalConBuffer <- 4
	}()

	num := <-canalConBuffer
	fmt.Println(num)

	time.Sleep(1 * time.Second)

	num = <-canalConBuffer
	fmt.Println(num)

	num = <-canalConBuffer
	fmt.Println(num)

	num = <-canalConBuffer
	fmt.Println(num)

	canalTareas := make(chan string)

	tareas := []string{
		"Ir a comprar sandia",
		"Pagar el seguro",
		"Llamar a la abuela",
		"Recoger las llaves de la moto",
	}

	// go func() {
	// 	for _, tarea := range tareas {
	// 		canalTareas <- tarea
	// 	}
	// 	close(canalTareas)
	// }()
	var wcTareas sync.WaitGroup
	wcTareas.Add(2)

	go realizarTareas(tareas, canalTareas, &wcTareas)

	go recibirTareasHechas(canalTareas, &wcTareas)

	// valor, ok := <-canalTareas
	// if !ok {
	// 	fmt.Println("El canal está cerrado.")
	// } else {
	// 	fmt.Println("El canal está abierto: ", valor)
	// }

	// for tarea := range canalTareas {
	// 	fmt.Printf("%s: realizada\n", tarea)
	// }

	// valor, ok = <-canalTareas
	// if !ok {
	// 	fmt.Println("El canal está cerrado.")
	// } else {
	// 	fmt.Println("El canal está abierto: ", valor)
	// }

	wcTareas.Wait()

	fmt.Println("Ya no quedan mas tareas por hacer")

	canalApi1 := make(chan string)
	canalApi2 := make(chan string)

	go getDatosApi1(canalApi1)
	go getDatosApi2(canalApi2)

	canal1Cerrado := false
	canal2Cerrado := false

	for !canal1Cerrado || !canal2Cerrado {
		select {
		case resp1 := <-canalApi1:
			fmt.Println(resp1)
			canal1Cerrado = true
			canalApi1 = nil
		case resp2 := <-canalApi2:
			fmt.Println(resp2)
			canal2Cerrado = true
			canalApi2 = nil
		}
	}

	fmt.Println("Ya han contestado las APIs")

	// Ejercicio:
	// Receta -> []Paso, Nombre
	// Paso -> Duracion, Descripcion
	// Método hacerReceta -> iterar los pasos y va a ir avisando por un canal el progreso
	// Al finalizar las 2 recetas, avisar que ya pueden llevarlas a la mesa

}
