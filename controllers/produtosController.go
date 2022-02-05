package controllers

import (
	"daniel/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
    if r.Method == "Post" {
        nome := r.FormValue("nome")
        descricao := r.FormValue("descricao")
        preco := r.FormValue("preco")
        quantidade := r.FormValue("quantidade")

        precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
        if err != nil {
            log.Println("Erro ao converter preço")
        }

        quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
        if err != nil {
            log.Println("Erro ao converter quantidade")
        }

        models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
    }
    http.Redirect(w, r, "/", 301)
}
