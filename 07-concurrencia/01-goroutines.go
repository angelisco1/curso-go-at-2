package main

import (
	"fmt"
	"sync"
)

func mostrarTexto(texto string, wg *sync.WaitGroup) {
	defer wg.Done()

	// time.Sleep(5 * time.Second)
	fmt.Println(texto)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	go mostrarTexto("Paso 1", &wg)
	go mostrarTexto("Paso 2", &wg)
	go mostrarTexto("Paso 3", &wg)

	go mostrarTexto("Fin", &wg)

	wg.Wait()

	fmt.Println("Fin")
	// time.Sleep(1 * time.Second)
}
