package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
	}
	servidor := &ServidorJogador{&armazenamento}

	t.Run("retorna pontuacao de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Maria")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarCorpoRequisicao(t, resposta.Body.String(), "20")
	})

	t.Run("retorna pontuacao de Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarCorpoRequisicao(t, resposta.Body.String(), "10")
	})
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisição é inválido, obtive '%s' esperava '%s'", recebido, esperado)
	}
}

func TestMeuCrud(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/account", nil)
		resposta := httptest.NewRecorder()

		os.Setenv("Pass", "<yourPass>")
		store, err := NewPostgresStore()
		if err != nil {
			log.Fatal(err)
		}

		if err := store.Init(); err != nil {
			log.Fatal(err)
		}
		server := NewAPIServer(":3000", store)
		//server.Run()
		get := server.handleAccount(resposta, requisicao)
		verificarCorpoRequisicao(t, resposta.Body.String(), "[{\"id\":8,\"firstName\":\"Matheus100\",\"lastName\":\"Vasconcelos100\",\"number\":100,\"balance\":0,\"createdAt\":\"2023-01-14T19:42:24.510858Z\"},{\"id\":9,\"firstName\":\"Matheus101\",\"lastName\":\"Vasconcelos101\",\"number\":101,\"balance\":0,\"createdAt\":\"2023-01-14T19:44:01.814875Z\"},{\"id\":10,\"firstName\":\"Matheus103\",\"lastName\":\"Vasconcelos103\",\"number\":498081,\"balance\":0,\"createdAt\":\"2023-01-14T19:46:07.30702Z\"},{\"id\":11,\"firstName\":\"Matheus104\",\"lastName\":\"Vasconcelos104\",\"number\":727887,\"balance\":0,\"createdAt\":\"2023-01-14T19:46:15.490114Z\"},{\"id\":12,\"firstName\":\"Matheus105\",\"lastName\":\"Vasconcelos105\",\"number\":131847,\"balance\":0,\"createdAt\":\"2023-01-14T19:46:25.712571Z\"},{\"id\":13,\"firstName\":\"Matheus106\",\"lastName\":\"Vasconcelos106\",\"number\":984059,\"balance\":0,\"createdAt\":\"2023-01-14T19:46:37.454121Z\"}]\n")
		fmt.Println(get)

	})
	t.Run("post", func(t *testing.T) {
		values := map[string]string{"firstName": "foo2", "lastName": "baz2"}
		jsonData, err := json.Marshal(values)
		if err != nil {
			log.Fatal(err)
		}
		requisicao, _ := http.NewRequest(http.MethodPost, "/account", bytes.NewBuffer(jsonData))
		resposta := httptest.NewRecorder()

		os.Setenv("Pass", "<yourPass>")
		store, err := NewPostgresStore()
		if err != nil {
			log.Fatal(err)
		}

		if err := store.Init(); err != nil {
			log.Fatal(err)
		}
		server := NewAPIServer(":3000", store)
		//server.Run()
		server.handleAccount(resposta, requisicao)

		if resposta.Code != 200 {
			t.Errorf("Status esperado: '%o' recebido '%o'", 200, resposta.Code)
		}

	})
}

type EsbocoArmazenamentoJogador struct {
	pontuacoes map[string]int
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}
