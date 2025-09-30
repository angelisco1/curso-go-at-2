package bbdd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var BBDD *sql.DB

func ConectarBBDD() {
	BBDD, err := sql.Open("sqlite3", "./bbdd.db")
	if err != nil {
		log.Fatalf("Error al conectarnos a la BBDD: %v", err)
	}

	sql := `
	CREATE TABLE IF NOT EXISTS acciones (
		id TEXT PRIMARY KEY NOT NULL,
		ticker TEXT NOT NULL UNIQUE,
		nombre TEXT NOT NULL,
		precio REAL NOT NULL
	);
	`

	_, err = BBDD.Exec(sql)
	if err != nil {
		log.Fatalf("Error al crear la tabla: %v", err)
	}

	fmt.Println("Todo ok hasta ahora")
}
