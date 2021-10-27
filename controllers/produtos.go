package controllers

import (
	"html/template"
	"log"
	"net/http"
	"soder_loja/models"
	"strconv"
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

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro de converão do preço", err)
		}

		qntCnvtInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro de converão do quantidade", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, qntCnvtInt)
		http.Redirect(w, r, "/", 301)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	produto := models.EditaProduto(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertida, err := strconv.Atoi(id)

		if err != nil {
			log.Panicln("Erro na convers'ao para int")
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Panicln("Erro na convers'ao para preco")
		}
		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Panicln("Erro na convers'ao para quantidade")
		}
		models.AtualizaProduto(idConvertida, nome, precoConvertido, quantidadeConvertido, descricao)
	}
	http.Redirect(w, r, "/", 301)
}
