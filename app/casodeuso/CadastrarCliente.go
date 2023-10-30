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

func (uc *CadastrarCliente) CadastrarCliente(inputCliente apresentacao.ClienteDTO) error {
	hash := md5.Sum([]byte(inputCliente.CPF))

	cliente := dominio.NewCliente(hex.EncodeToString(hash[:]),
		inputCliente.CPF,
		inputCliente.Nome,
		inputCliente.Email,
	)

	err := uc.clienteRepository.CadastrarCliente(cliente)
	if err != nil {
		return err
	}

	return nil
}
