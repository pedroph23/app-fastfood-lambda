package controladores

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type AutenticacaoController struct {
	consultarClienteUC    usecases.ConsultarClienteUC
	autenticacaoClienteUC usecases.AutenticacaoClienteUC
}

func NewAutenticacaoController(consultarClienteUC usecases.ConsultarClienteUC, autenticacaoClienteUC usecases.AutenticacaoClienteUC) *AutenticacaoClienteController {
	return &AutenticacaoController{
		consultarClienteUC:    consultarClienteUC,
		autenticacaoClienteUC: autenticacaoClienteUC,
	}
}

func (c *AutenticacaoController) Handle(idCliente string) ([]byte, error) {
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
