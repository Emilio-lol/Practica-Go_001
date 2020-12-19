package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err    error
	db     *sql.DB
	nombre string
)

func main() {
	nombre = "Emilio Andere Lopez"
	AbrirConexion()
	data := ShowQuery("membresia", nombre)
	if data != "0" {
		fmt.Println(data)
	} else {
		dato, id := InsertQuery(nombre)
		if dato {
			fmt.Println("Insertado con exito\n" + nombre + ". ID: " + id)
		} else {
			fmt.Println("A ocurrido un error")
		}
	}
}

func AbrirConexion() {
	db, err = sql.Open("mysql", string("root:@tcp(127.0.0.1:3306)/test_pract01"))
	if err != nil {
		panic(err)

	}
	err = db.Ping()
	if err != nil {
		panic(err)

	}
}

func ShowQuery(tabla string, where string) string {
	var id, tipo string
	query := db.QueryRow("SELECT * FROM "+tabla+" WHERE tipo_membresia = ? ", where)
	err = query.Scan(&id, &tipo)
	if err != nil {
		return "0"
	}
	return tipo + ". ID: " + id
}

func InsertQuery(Name string) (bool, string) {
	add, err := db.Exec("Insert into membresia (id_membresia, tipo_membresia) values (NULL, ?)", Name)
	if err != nil {
		return false, "0"
	}
	query, err := add.LastInsertId()
	if err != nil {
		return false, "0"
	}
	return true, strconv.Itoa(int(query))
}
