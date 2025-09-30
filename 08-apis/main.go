package main

import (
	"apis/bbdd"
	"apis/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	bbdd.ConectarBBDD()
	defer bbdd.BBDD.Close()

	router := mux.NewRouter()

	router.HandleFunc("/acciones", controllers.GetAcciones).Methods("GET")
	router.HandleFunc("/acciones", controllers.PostAcciones).Methods("POST")
	router.HandleFunc("/acciones/{id}", controllers.GetAccion).Methods("GET")

	puerto := "3000"
	fmt.Printf("Escuchando en http://localhost:%s\n", puerto)

	err := http.ListenAndServe(":"+puerto, router)
	if err != nil {
		log.Fatalf("Fallo al arrancar el servidor...")
	}
}
