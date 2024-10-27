package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	http.HandleFunc("/", CepHandler)
	http.ListenAndServe(":8080", nil)
}

func CepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

// cep parametro de entrada,   (ViaCEP, error) retorno
func BuscaCep(cep string) (*ViaCep, error) {
	// fazer uma request web com o arq lido
	req, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json", cep))
	if err != nil {
		return nil, err
	}

	// ler a resposta
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	var c ViaCep
	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
