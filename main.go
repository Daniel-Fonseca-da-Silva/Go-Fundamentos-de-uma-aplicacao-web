package main

import (
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Relógio ourax", Descricao: "Super moderno e elegante", Preco: 13299.99, Quantidade: 15},
		{"Teclado Jumpo", "Teclado por ondas radionicas", 59.99, 851},
		{"Video game grudano", "Super video game perca vida real", 89.99, 551},
		{"Fonte potente nanomax", "Super fonte para nave espacial", 885.99, 125},
		{"Espada alienigena da era antiga", "Espada encontrada no planeta zapio", 1359.99, 2},
		{"Esqueleto de humanoides 2022", "Esqueleto humanoide usado para hologramar", 359.99, 155},
		{"Pistola de dupla materia 2055", "Super pistola com disparo quadriplo ou triplo", 3285.74, 78},
	}
	temp.ExecuteTemplate(w, "Index", produtos)

}
