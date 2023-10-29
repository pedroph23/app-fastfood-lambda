// repositories/cliente_repository.go
package repositorio

import (
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type RepositorioCliente interface {
	CadastrarCliente(cliente *dominio.Cliente) error
	BuscarClientePorID(idCliente string) (*dominio.Cliente, error)
}
