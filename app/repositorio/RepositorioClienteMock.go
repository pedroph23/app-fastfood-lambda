package repositorio

import (
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

// RepositorioClienteMock é um mock do repositório do cliente.
type RepositorioClienteMock struct{}

// NewRepositorioClienteMock cria uma nova instância de RepositorioClienteMock.
func NewRepositorioClienteMock() *RepositorioClienteMock {
	return &RepositorioClienteMock{}
}

// SalvarOuAtualizarCliente simula a operação de salvar ou atualizar um cliente.
func (r *RepositorioClienteMock) SalvarOuAtualizarCliente(cliente *dominio.Cliente) error {
	// Aqui você pode retornar um cliente mockado
	return nil
}

// BuscarClientePorID simula a operação de buscar um cliente pelo ID.
func (r *RepositorioClienteMock) BuscarClientePorID(idCliente string) (*dominio.Cliente, error) {
	// Aqui você pode retornar um cliente mockado
	return &dominio.Cliente{
		ID:     "123",
		CPF:    "12345678900",
		Nome:   "John Doe",
		Email:  "john.doe@example.com",
		Status: "ATIVO",
	}, nil
}
