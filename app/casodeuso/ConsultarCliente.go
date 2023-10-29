// usecases/consultar_cliente_uc.go
package usecases

import (
	"fmt"

	"github.com/example/domain"
	"github.com/example/repositories"
)

type ConsultarClienteUC struct {
	clienteRepository repositories.ClienteRepository
}

func NewConsultarClienteUC(clienteRepository repositories.ClienteRepository) *ConsultarClienteUC {
	return &ConsultarClienteUC{clienteRepository: clienteRepository}
}

func (uc *ConsultarClienteUC) ConsultarCliente(idCliente string) (*domain.Cliente, error) {
	cliente, err := uc.clienteRepository.BuscarClientePorID(idCliente)
	if err != nil {
		return nil, fmt.Errorf("failed to find client: %v", err)
	}

	domainCliente := domain.NewCliente(cliente.CPF, cliente.Nome, cliente.Email, cliente.ID)

	return domainCliente, nil
}
