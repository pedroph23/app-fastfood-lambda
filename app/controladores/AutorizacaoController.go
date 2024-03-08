package controladores

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
)

type AutorizarcaoController struct {
	consultarClienteUC casodeuso.ConsultarCliente
	autorizarClienteUC casodeuso.AutorizarUsuario
}

func NewAutorizarcaoController(consultarClienteUC casodeuso.ConsultarCliente, autorizarClienteUC casodeuso.AutorizarUsuario) *AutorizarcaoController {
	return &AutorizarcaoController{
		consultarClienteUC: consultarClienteUC,
		autorizarClienteUC: autorizarClienteUC,
	}
}

func (c *AutorizarcaoController) Handle(tokenString string, methodArn string) (events.APIGatewayCustomAuthorizerResponse, error) {

	ok, idUsuario := c.autorizarClienteUC.AutorizarCliente(tokenString)
	fmt.Printf("idUsuario: %s\n", idUsuario)
	fmt.Println("ok: ", ok)
	if !ok {
		// Se não existir token, libere a requisição
		return events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: "anonymous",
			PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   "Deny",
						Resource: []string{methodArn},
					},
				},
			},
		}, nil
	} else if idUsuario == "" {
		// Usuario anonimo
		return events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: "anonymous",
			PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   "Allow",
						Resource: []string{methodArn},
					},
				},
			},
		}, nil
	}

	_, err := c.consultarClienteUC.ConsultarCliente(idUsuario)
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, err
	}

	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: idUsuario,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{methodArn},
				},
			},
		},
	}, nil

}
