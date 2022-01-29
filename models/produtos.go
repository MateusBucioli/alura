package models

import (
	"alura/db"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProdutos() []Produto {
	db := db.DBConn()

	selectProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic("Falha ao consultar produtos. Motivo: " + err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic("Falha ao consultar produtos. Motivo: " + err.Error())
		}

		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	defer db.Close()

	return produtos
}

func SetProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.DBConn()

	insert, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
