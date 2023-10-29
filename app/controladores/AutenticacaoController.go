package controladores

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
)

type AutenticacaoController struct {
	consultarClienteUC    *casodeuso.ConsultarCliente
	autenticacaoClienteUC *casodeuso.AutenticarUsuario
}

func NewAutenticacaoController(consultarClienteUC *casodeuso.ConsultarCliente, autenticacaoClienteUC *casodeuso.AutenticarUsuario) *AutenticacaoController {
	return &AutenticacaoController{
		consultarClienteUC:    consultarClienteUC,
		autenticacaoClienteUC: autenticacaoClienteUC,
	}
}

func (c *AutenticacaoController) Handle(idCliente string) ([]byte, error) {
	hash := md5.Sum([]byte(idCliente))
	hashStr := hex.EncodeToString(hash[:])
	println("hashStr")
	println(hashStr)

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
