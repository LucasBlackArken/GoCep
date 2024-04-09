package application

import (
	"github.com/LucasBlackArken/GoCep/domain"
	"github.com/LucasBlackArken/GoCep/infrastructure"
)

type EnderecoService struct {
	viaCEP *infrastructure.ViaCEP
}

func NovoEndereco(viaCEP *infrastructure.ViaCEP) *EnderecoService {
	return &EnderecoService{
		viaCEP: viaCEP,
	}
}

func (s *EnderecoService) BuscaEndereco(cep string) (*domain.Endereco, error) {
	return s.viaCEP.BuscaEndereco(cep)
}
