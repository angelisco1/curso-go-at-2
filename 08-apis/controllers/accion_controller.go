package controllers

import (
	"apis/bbdd"
	"apis/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var Acciones []models.Accion = []models.Accion{
	models.Accion{"abc", "TSLA", "Tesla", 440.62},
	models.Accion{"def", "AMZN", "Amazon", 240.62},
	models.Accion{"ghi", "APPL", "Apple", 340.62},
}

func GetAcciones(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get acciones")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Acciones)
}

func PostAcciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Post acciones")

	var accionCreada models.Accion

	json.NewDecoder(r.Body).Decode(&accionCreada)

	accionCreada.Id = uuid.NewString()
	fmt.Println(accionCreada)

	_, err := bbdd.BBDD.Exec("INSERT INTO acciones (id, ticker, nombre, precio) VALUES (?, ?, ?, ?)", accionCreada.Id, accionCreada.Ticker, accionCreada.Nombre, accionCreada.Precio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear la accion: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(accionCreada)
}

func GetAccion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get acción")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println("->", id)

	var accionBuscada models.Accion

	result := bbdd.BBDD.QueryRow("SElECT id, ticker, nombre, precio FROM acciones WHERE id = ?", id)
	err := result.Scan(&accionBuscada.Id, &accionBuscada.Ticker, &accionBuscada.Nombre, &accionBuscada.Precio)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("Accion con id %s no encontrada", id), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error al buscar en la BBDD: %v", err), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accionBuscada)
}
