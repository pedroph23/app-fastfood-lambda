package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

type Response struct {
	Message string `json:"message"`
}

func AutenticacaoClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest, autenticacaoClienteUC *casodeuso.AutenticarUsuario,
	consultarClienteUC *casodeuso.ConsultarCliente) (events.APIGatewayProxyResponse, error) {
	// TODO: Implementar a lógica de autenticação do cliente
	controller := controladores.NewAutenticacaoController(consultarClienteUC, autenticacaoClienteUC)
	respBody, err := controller.Handle(req.PathParameters["id_cliente"])
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func CadastroClienteHandler(ctx context.Context, req events.APIGatewayProxyRequest, cadastrarClienteUC *casodeuso.CadastrarCliente) (events.APIGatewayProxyResponse, error) {
	// TODO: Implementar a lógica de criação de cliente
	controller := controladores.NewCadastroClienteController(cadastrarClienteUC)
	log.Printf("req.Body: %s\n", req.Body)
	respBody, err := controller.Handle(req.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest, autenticacaoClienteUC *casodeuso.AutenticarUsuario,
	consultarClienteUC *casodeuso.ConsultarCliente, cadastrarClienteUC *casodeuso.CadastrarCliente) (events.APIGatewayProxyResponse, error) {
	log.Printf("req.Body: %s\n", req.Body)
	switch req.HTTPMethod {
	case "POST":
		if req.Path == "/clientes/{id_cliente}/auth" {
			return AutenticacaoClienteHandler(ctx, req, autenticacaoClienteUC, consultarClienteUC)
		} else if req.Path == "/clientes" {
			return CadastroClienteHandler(ctx, req, cadastrarClienteUC)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}, nil
}

func main() {
	clienteRepository := repositorio.NewRepositorioClienteImpl()
	autenticacaoClienteUC := casodeuso.NewAutenticarUsuario()
	consultarClienteUC := casodeuso.NewConsultarCliente(clienteRepository)
	cadastrarClienteUC := casodeuso.NewCadastrarCliente(clienteRepository)

	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Printf("req.Body: %s\n", req.Body)
		return Handler(ctx, req, autenticacaoClienteUC, consultarClienteUC, cadastrarClienteUC)

	})
}
