package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/app-fastfood-lambda/app/casodeuso"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/example/controllers"
	"github.com/example/repositories"
)

type Response struct {
	Message string `json:"message"`
}

func AutenticacaoClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest, autenticacaoClienteUC usecases.AutenticacaoClienteUC) (events.APIGatewayProxyResponse, error) {
	// TODO: Implementar a lógica de autenticação do cliente
	controller := controllers.NewAutenticacaoClienteController(autenticacaoClienteUC)
	respBody, err := controller.Handle(req.PathParameters["id_cliente"])
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func CadastroClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest, cadastroClienteUC usecases.CadastroClienteUC) (events.APIGatewayProxyResponse, error) {
	// TODO: Implementar a lógica de criação de cliente
	controller := controllers.NewCadastroClienteController(cadastroClienteUC)
	respBody, err := controller.Handle(req.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest, autenticacaoClienteUC usecases.AutenticacaoClienteUC, cadastroClienteUC usecases.CadastroClienteUC) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		if req.Path == "/clientes/{id_cliente}/auth" {
			return AutenticacaoClienteHandler(ctx, req, autenticacaoClienteUC)
		} else if req.Path == "/clientes" {
			return CadastroClienteHandler(ctx, req, cadastroClienteUC)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}, nil
}

func main() {
	clienteRepository := repositories.NewClienteRepository()
	autenticacaoClienteUC := usecases.NewAutenticacaoClienteUC(clienteRepository)
	cadastroClienteUC := usecases.NewCadastroClienteUC(clienteRepository)

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return Handler(ctx, req, autenticacaoClienteUC, cadastroClienteUC)
	})
}
