// repositories/cliente_repository.go
package repositorio

import "github.com/example/domain"

type RepositorioCliente interface {
	CadastrarCliente(cliente *domain.Cliente) error
	BuscarClientePorID(idCliente string) (*domain.Cliente, error)
}
