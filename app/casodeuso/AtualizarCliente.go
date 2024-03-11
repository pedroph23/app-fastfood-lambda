package casodeuso

import (
	"github.com/pedroph23/app-fastfood-lambda/app/apresentacao"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type AtualizarCliente interface {
	AtualizarCliente(inputCliente *dominio.Cliente, novosDadosCliente *apresentacao.ClienteDTO) (*dominio.Cliente, error)
}
