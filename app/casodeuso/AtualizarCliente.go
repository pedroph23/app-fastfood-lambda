package casodeuso

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

type AtualizarCliente struct {
	clienteRepository repositorio.RepositorioCliente
}

func NewAtualizarCliente(clienteRepository repositorio.RepositorioCliente) *AtualizarCliente {
	return &AtualizarCliente{
		clienteRepository: clienteRepository,
	}
}

func (uc *AtualizarCliente) AtualizarCliente(inputCliente *dominio.Cliente) (*dominio.Cliente, error) {
	var cliente *dominio.Cliente
	var err error
	if inputCliente.Status == "INATIVO" {
		cliente, err = dominio.NewCliente(
			"",
			inputCliente.ID,
			"",
			"",
			"INATIVO",
		)
	} else {
		hash := md5.Sum([]byte(inputCliente.CPF))
		cliente, err = dominio.NewCliente(
			inputCliente.CPF,
			hex.EncodeToString(hash[:]),
			inputCliente.Nome,
			inputCliente.Email,
			"ATIVO",
		)
	}

	if err != nil {
		return nil, err
	}

	err = uc.clienteRepository.SalvarOuAtualizarCliente(cliente)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}
