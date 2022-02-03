package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DBConn() *sql.DB {
	conn := "user=postgres dbname=alura password=!@#PSQL host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	return db
}
