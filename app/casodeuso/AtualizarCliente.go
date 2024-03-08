package casodeuso

import (
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type AtualizarCliente interface {
	AtualizarCliente(inputCliente *dominio.Cliente) (*dominio.Cliente, error)
}
