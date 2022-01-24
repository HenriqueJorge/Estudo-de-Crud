package main

import (
	"crudGo/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Estrutura basica do crud
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cria", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/busca", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/atualiza", servidor.AtualizaUsuario).Methods(http.MethodPut)
	router.HandleFunc("/deleta", servidor.DeletaUsuario).Methods(http.MethodDelete)

	fmt.Println("Escultando porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
