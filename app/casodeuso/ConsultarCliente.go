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

	domainCliente, err := dominio.NewCliente(cliente.CPF, cliente.ID, cliente.Nome, cliente.Email)

	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	return domainCliente, nil
}
