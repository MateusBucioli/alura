package controllers

import (
	"alura/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preco:", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade:", err)
		}

		models.SetProduto(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", 301)
}
