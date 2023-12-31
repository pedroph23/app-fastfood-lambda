// casodeuso/CadastrarCliente.go
package casodeuso

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/pedroph23/app-fastfood-lambda/app/apresentacao"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

type CadastrarCliente struct {
	clienteRepository repositorio.RepositorioCliente
}

func NewCadastrarCliente(clienteRepository repositorio.RepositorioCliente) *CadastrarCliente {
	return &CadastrarCliente{
		clienteRepository: clienteRepository,
	}
}

func (uc *CadastrarCliente) CadastrarCliente(inputCliente apresentacao.ClienteDTO) (*dominio.Cliente, error) {
	hash := md5.Sum([]byte(inputCliente.CPF))

	cliente, err := dominio.NewCliente(
		inputCliente.CPF,
		hex.EncodeToString(hash[:]),
		inputCliente.Nome,
		inputCliente.Email,
	)

	if err != nil {
		return nil, err
	}

	err = uc.clienteRepository.CadastrarCliente(cliente)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}
