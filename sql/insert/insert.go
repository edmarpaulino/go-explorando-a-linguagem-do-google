package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// statement - substitui o ? pelo valor passado como argumento no método Exec
	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	rows, _ := res.RowsAffected()
	fmt.Println(rows)
}
