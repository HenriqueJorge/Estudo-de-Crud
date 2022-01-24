package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler arquivo"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro na conversao"))
		return
	}

	fmt.Println(usuario)

	w.Write([]byte("Usuario criado"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca usuario"))
	return
}

func AtualizaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza usuario"))
	return
}

func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete usuario"))
	return
}
