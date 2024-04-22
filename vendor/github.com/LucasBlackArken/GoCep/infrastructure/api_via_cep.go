package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LucasBlackArken/GoCep/domain"
)

type ViaCEP struct{}

func (v *ViaCEP) BuscaEndereco(cep string) (*domain.Endereco, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	response, erro := http.Get(url)
	if erro != nil {
		return nil, erro
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erro ao buscar endereço. Status: %s", response.Status)
	}

	var endereco domain.Endereco
	erro = json.NewDecoder(response.Body).Decode(&endereco)
	if erro != nil {
		return nil, erro
	}

	return &endereco, nil
}
