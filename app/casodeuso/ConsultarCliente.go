package casodeuso

import (
	"fmt"

	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

type ConsultarCliente struct {
	clienteRepository repositorio.RepositorioCliente
}

func NewConsultarCliente(clienteRepository repositorio.RepositorioCliente) *ConsultarCliente {
	return &ConsultarCliente{clienteRepository: clienteRepository}
}

func (uc *ConsultarCliente) ConsultarCliente(idCliente string) (*dominio.Cliente, error) {
	cliente, err := uc.clienteRepository.BuscarClientePorID(idCliente)
	if err != nil {
		return nil, fmt.Errorf("failed to find client: %v", err)
	}

	domainCliente := dominio.NewCliente(cliente.CPF, cliente.Nome, cliente.Email, cliente.ID)

	return domainCliente, nil
}
