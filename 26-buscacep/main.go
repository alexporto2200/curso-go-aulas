package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	// ler do terminal o arg1
	for _, cep := range os.Args[1:] {

		// fazer uma request web com o arq lido
		req, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro request: %v\n", err)
		}

		// ler a resposta
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error read response %v\n", err)
		}

		// converter a resposta para json
		var data ViaCep

		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error convert json %v\n", err)
		}

		fmt.Println(data)

		// escrever a resposta num arquivo
		file, err := os.Create("arquivo.txt")

		if err != nil {
			panic(err)

		}

		file.WriteString(fmt.Sprintf("Cep: %s, Endereço %s", data.Cep, data.Localidade))

		defer file.Close()

		// ler o arquivo e imprimir
		rFile, err := os.ReadFile("arquivo.txt")

		if err != nil {
			panic(err)
		}

		fmt.Println(string(rFile))
	}
}
