package controladores

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
)

type AutorizarcaoController struct {
	consultarClienteUC *casodeuso.ConsultarCliente
	autorizarClienteUC *casodeuso.AutorizarUsuario
}

func NewAutorizarcaoController(consultarClienteUC *casodeuso.ConsultarCliente, autorizarClienteUC *casodeuso.AutorizarUsuario) *AutorizarcaoController {
	return &AutorizarcaoController{
		consultarClienteUC: consultarClienteUC,
		autorizarClienteUC: autorizarClienteUC,
	}
}

func (c *AutorizarcaoController) Handle(tokenString string) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {

	ok, idUsuario := c.autorizarClienteUC.AutorizarCliente(tokenString)
	fmt.Printf("idUsuario: %s\n", idUsuario)
	fmt.Println("ok: ", ok)
	if !ok {
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
		}, nil
	} else {
		if idUsuario == "" {
			return events.APIGatewayV2CustomAuthorizerSimpleResponse{
				IsAuthorized: true,
			}, nil
		}
	}

	_, err := c.consultarClienteUC.ConsultarCliente(idUsuario)
	if err != nil {
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{}, err
	}

	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
	}, nil
}
