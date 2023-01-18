package main

import (
	"log"
	"net/http"
)

type ArmazenamentoJogadorEmMemoria struct{}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return 123
}

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	//server := NewAPIServer(":3000", store)
	//server.Run()

	servidor := &ServidorJogador{&ArmazenamentoJogadorEmMemoria{}}

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}
