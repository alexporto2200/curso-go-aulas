package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 1, Saldo: 2}

	// transforma a strutura para json
	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	println(string(res))

	// faz o encode do conta, e joga na saida padrao stdout
	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		panic(err)
	}

	// string de json em strutura

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 3}`)

	var contaC Conta

	err = json.Unmarshal(jsonPuro, &contaC)
	if err != nil {
		panic(err)
	}

	println(contaC.Numero, contaC.Saldo)

	// string de json para estrutura com outro nome

	type ContaJ struct {
		Numero int `json:"n"`
		Saldo  int `json:"s"`
	}

	jsonPuroJ := []byte(`{"n": 5, "s": 6}`)

	var outraConta ContaJ

	err = json.Unmarshal(jsonPuroJ, &outraConta)

	if err != nil {
		panic(err)
	}

	println(outraConta.Numero, outraConta.Saldo)

}
