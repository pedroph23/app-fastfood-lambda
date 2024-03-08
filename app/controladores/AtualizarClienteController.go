package controladores

import (
	"encoding/json"
	"fmt"

	"github.com/pedroph23/app-fastfood-lambda/app/apresentacao"
	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type AtualizarClienteController struct {
	atualizarClienteUC casodeuso.AtualizarCliente
	consultaClienteUC  casodeuso.ConsultarCliente
}

func NewAtualizarClienteController(atualizarClienteUC casodeuso.AtualizarCliente,
	consultaClienteUC casodeuso.ConsultarCliente) *AtualizarClienteController {
	return &AtualizarClienteController{
		atualizarClienteUC: atualizarClienteUC,
		consultaClienteUC:  consultaClienteUC,
	}
}

func (controller *AtualizarClienteController) Handle(idCliente string, requestBody string) ([]byte, error) {
	var clienteDTO apresentacao.ClienteDTO
	var cliente *dominio.Cliente
	fmt.Printf("req.Body: %s\n", requestBody)

	err := json.Unmarshal([]byte(requestBody), &clienteDTO)

	if err != nil {
		return nil, err
	}

	cliente, err = controller.consultaClienteUC.ConsultarCliente(clienteDTO.ID)

	if err != nil {
		return nil, err
	}

	cliente, err = controller.atualizarClienteUC.AtualizarCliente(cliente)

	if err != nil {
		return nil, err
	}

	response := map[string]string{
		"message":    "Cliente atualizado com sucesso",
		"id_cliente": cliente.ID,
	}
	respBody, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
