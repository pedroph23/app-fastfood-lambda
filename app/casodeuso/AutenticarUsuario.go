package casodeuso

import (
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type AutenticarUsuario interface {
	AutenticarClienteAnonimo() (string, error)
	AutenticarCliente(cliente *dominio.Cliente) (string, error)
}
