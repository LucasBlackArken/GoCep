package main

import (
	"fmt"

	"github.com/LucasBlackArken/GoCep/application"
	"github.com/LucasBlackArken/GoCep/infrastructure"
)

func main() {
	viaCEP := &infrastructure.ViaCEP{}
	enderecoService := application.NovoEndereco(viaCEP)

	cep := "01001000" // CEP de exemplo, você pode alterar para o CEP desejado

	endereco, erro := enderecoService.BuscaEndereco(cep)
	if endereco != nil {
		fmt.Println("Erro ao obter o endereço:", erro)
		return
	}

	fmt.Println("Endereço encontrado:")
	fmt.Println("CEP:", endereco.Cep)
	fmt.Println("Logradouro:", endereco.Logradouro)
	fmt.Println("Complemento:", endereco.Complemento)
	fmt.Println("Bairro:", endereco.Bairro)
	fmt.Println("Cidade:", endereco.Localidade)
	fmt.Println("UF:", endereco.Uf)
}
