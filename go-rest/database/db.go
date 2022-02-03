package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnect() {
	connStr := "host=localhost user=postgres password=postgres dbname=postgres port=5444"

	fmt.Println("Conectando com o banco...")
	start := time.Now()

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	elapsed := time.Since(start)
	fmt.Println("Conexão concluída em:", elapsed.String())

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
