package controladores

import (
	"fmt"

	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
)

type AutenticacaoController struct {
	consultarClienteUC    casodeuso.ConsultarCliente
	autenticacaoClienteUC casodeuso.AutenticarUsuario
}

func NewAutenticacaoController(consultarClienteUC casodeuso.ConsultarCliente, autenticacaoClienteUC casodeuso.AutenticarUsuario) *AutenticacaoController {
	return &AutenticacaoController{
		consultarClienteUC:    consultarClienteUC,
		autenticacaoClienteUC: autenticacaoClienteUC,
	}
}

func (c *AutenticacaoController) Handle(idCliente string) ([]byte, error) {
	var token string
	if idCliente == "anonimo" {
		token, _ = c.autenticacaoClienteUC.AutenticarClienteAnonimo()
	} else {
		cliente, err := c.consultarClienteUC.ConsultarCliente(idCliente)
		if err != nil {
			return nil, fmt.Errorf("failed to authenticate client: %v", err)
		}

		token, err = c.autenticacaoClienteUC.AutenticarCliente(cliente)
		if err != nil {
			return nil, fmt.Errorf("failed to authenticate client: %v", err)
		}
	}

	return []byte(token), nil
}
