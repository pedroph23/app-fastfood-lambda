// controllers/autenticacao_cliente_controller.go
package controladores

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/app/casodeuso"
)

type AutenticacaoClienteController struct {
	consultarClienteUC    usecases.ConsultarClienteUC
	autenticacaoClienteUC usecases.AutenticacaoClienteUC
}

func NewAutenticacaoClienteController(consultarClienteUC usecases.ConsultarClienteUC, autenticacaoClienteUC usecases.AutenticacaoClienteUC) *AutenticacaoClienteController {
	return &AutenticacaoClienteController{
		consultarClienteUC:    consultarClienteUC,
		autenticacaoClienteUC: autenticacaoClienteUC,
	}
}

func (c *AutenticacaoClienteController) Handle(idCliente string) ([]byte, error) {
	hash := md5.Sum([]byte(idCliente))
	hashStr := hex.EncodeToString(hash[:])

	cliente, err := c.consultarClienteUC.ConsultarCliente(hashStr)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate client: %v", err)
	}

	token, err := c.autenticacaoClienteUC.AutenticarCliente(cliente)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate client: %v", err)
	}

	return []byte(token), nil
}
