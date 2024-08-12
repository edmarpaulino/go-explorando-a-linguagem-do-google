package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//  Começa uma transação
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		// extorna a transação
		tx.Rollback()
		log.Fatal(err)
	}

	// confirma a transação
	tx.Commit()
}
