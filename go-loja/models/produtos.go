package models

import (
	"alura/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProdutos() []Produto {
	db := db.DBConn()

	selectProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
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

		produto.Id = id
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

func DeleteProduto(id string) {
	db := db.DBConn()

	delete, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduto(id string) Produto {
	db := db.DBConn()

	produto, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoUpdate := Produto{}

	for produto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoUpdate.Id = id
		produtoUpdate.Nome = nome
		produtoUpdate.Descricao = descricao
		produtoUpdate.Preco = preco
		produtoUpdate.Quantidade = quantidade
	}

	defer db.Close()
	return produtoUpdate
}

func UpdateProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.DBConn()

	produto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	produto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
