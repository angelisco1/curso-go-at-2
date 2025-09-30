package models

type Accion struct {
	Id     string  `json:"id"`
	Ticker string  `json:"ticker"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}
