package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err  error
	db   *sql.DB
	name = "Emilio Andere Lopez"
)

type Membresia struct {
	Id   string
	Tipo string
}

func main() {
	openConn()
	registros := existMembresia(name)
	if registros != "0" {
		fmt.Println(registros)
	} else {
		data := insertMembresia(name)
		fmt.Println("Ha sido insertado:", data)
	}
}

func openConn() {
	//bienhechor:Bienhechor_1234;@tcp(74.208.31.248:3306)/bienhechor
	db, err = sql.Open("mysql", string("bienhechor:Bienhechor_1234;@tcp(74.208.31.248:3306)/bienhechor"))
	reviewError(err)
	err = db.Ping()
	reviewError(err)
}

func reviewError(err error) {
	if err != nil {
		panic(err)
	}
}

func existMembresia(name string) (status string) {
	var registros []Membresia
	mostrar, err := db.Query("SELECT * FROM Membresia")
	reviewError(err)

	for mostrar.Next() {
		var id, nombre string
		err = mostrar.Scan(&id, &nombre)
		reviewError(err)

		registros = append(registros, Membresia{id, nombre})
	}

	for i := 0; i < len(registros); i++ {
		if registros[i].Tipo == name {
			status = registros[i].Tipo + ". ID: " + registros[i].Id
			return
		}
	}
	status = "0"
	return
}

func insertMembresia(name string) bool {
	add, err := db.Exec("Insert into Membresia (IdMembresia, TipoMembresia) values (NULL, ?)", name)
	if err != nil {
		return false
	}
	_, err = add.LastInsertId()
	if err != nil {
		return false
	}
	return true
}
