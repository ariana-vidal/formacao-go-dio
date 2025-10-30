package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Produto struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome"`
	Preco     float64   `json:"preco"`
	Categoria Categoria `json:"categoria"`
}

type Categoria struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

var produtos []Produto

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(produtos)
}

func GetProdutoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range produtos {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Produto{})
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var produto Produto
	_ = json.NewDecoder(r.Body).Decode(&produto)
	produto.ID, _ = strconv.Atoi(params["id"])
	produtos = append(produtos, produto)
	json.NewEncoder(w).Encode(produtos)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range produtos {
		if strconv.Itoa(item.ID) == params["id"] {
			produtos = append(produtos[:index], produtos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(produtos)
}

func main() {
	router := mux.NewRouter()

	produtos = append(produtos, Produto{ID: 1, Nome: "Notebook", Preco: 2500.00, Categoria: Categoria{ID: 1, Nome: "Eletrônicos"}})
	produtos = append(produtos, Produto{ID: 2, Nome: "Smartphone", Preco: 1500.00, Categoria: Categoria{ID: 1, Nome: "Eletrônicos"}})

	router.HandleFunc("/produtos", GetProdutos).Methods("GET")
	router.HandleFunc("/produtos/{id}", GetProdutoByID).Methods("GET")
	router.HandleFunc("/produtos/{id}", CreateProduto).Methods("POST")
	router.HandleFunc("/produtos/{id}", DeleteProduto).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

