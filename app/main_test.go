package main

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/apresentacao"
	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

func TestHandler(t *testing.T) {
	// Mock do repositório do cliente para simular o banco de dados
	clienteRepository := repositorio.NewRepositorioClienteMock()

	// Instância dos casos de uso
	autenticacaoClienteUC := casodeuso.NewAutenticarUsuario()
	consultarClienteUC := casodeuso.NewConsultarCliente(clienteRepository)
	cadastrarClienteUC := casodeuso.NewCadastrarCliente(clienteRepository)
	atualizarClienteUC := casodeuso.NewAtualizarCliente(clienteRepository)
	autorizarUsuarioUC := casodeuso.NewAutorizarUsuario()

	t.Run("AutenticacaoClienteHandler", func(t *testing.T) {
		requestBody := `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "ATIVO"}`
		expectedResponse := "{\"access_token\":\"mocked_access_token\"}"

		req := events.APIGatewayProxyRequest{
			HTTPMethod:     "POST",
			Path:           "/auth",
			PathParameters: map[string]string{"id_cliente": "123"},
			Body:           requestBody,
		}

		response, err := controladores.AutenticacaoClienteHandler(context.Background(), req, autenticacaoClienteUC, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var authDTO apresentacao.AuthDTO
		err = json.Unmarshal([]byte(response.Body), &authDTO)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response.Body)
	})

	t.Run("CadastroClienteHandler", func(t *testing.T) {
		requestBody := `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "ATIVO"}`
		expectedResponse := "{\"message\":\"Cliente cadastrado com sucesso\",\"id_cliente\":\"123\"}"

		req := events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/clientes",
			Body:       requestBody,
		}

		response, err := controladores.CadastroClienteHandler(context.Background(), req, cadastrarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedResponse, response.Body)
	})

	t.Run("ConsultaClienteHandler", func(t *testing.T) {
		expectedID := "123"
		expectedCPF := "12345678900"
		expectedNome := "John Doe"
		expectedEmail := "john.doe@example.com"
		expectedStatus := "ATIVO"

		req := events.APIGatewayProxyRequest{
			HTTPMethod:     "GET",
			Path:           "/clientes/123",
			PathParameters: map[string]string{"id_cliente": "123"},
		}

		response, err := controladores.ConsultaClienteHandler(context.Background(), req, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var clienteDTO apresentacao.ClienteDTO
		err = json.Unmarshal([]byte(response.Body), &clienteDTO)
		assert.NoError(t, err)
		assert.Equal(t, expectedID, clienteDTO.ID)
		assert.Equal(t, expectedCPF, clienteDTO.CPF)
		assert.Equal(t, expectedNome, clienteDTO.Nome)
		assert.Equal(t, expectedEmail, clienteDTO.Email)
		assert.Equal(t, expectedStatus, clienteDTO.Status)
	})

	t.Run("AtualizaClienteHandler", func(t *testing.T) {
		requestBody := `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "INATIVO"}`
		expectedResponse := "{\"message\":\"Cliente atualizado com sucesso\",\"id_cliente\":\"123\"}"

		req := events.APIGatewayProxyRequest{
			HTTPMethod: "PATCH",
			Path:       "/clientes/123",
			PathParameters: map[string]string{
				"id_cliente": "123",
			},
			Body: requestBody,
		}

		response, err := controladores.AtualizaClienteHandler(context.Background(), req, atualizarClienteUC, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedResponse, response.Body)
	})

	// Teste de autorização
	t.Run("CustomAuthorizerHandler", func(t *testing.T) {
		token := "mocked_token"
		methodArn := "arn:aws:execute-api:region:account-id:api-id/stage/method/resource-path"
		expectedPrincipalID := "123"

		req := events.APIGatewayCustomAuthorizerRequest{
			AuthorizationToken: token,
			MethodArn:          methodArn,
		}

		response, err := controladores.CustomAuthorizerHandler(context.Background(), req, consultarClienteUC, autorizarUsuarioUC)
		assert.NoError(t, err)
		assert.Equal(t, expectedPrincipalID, response.PrincipalID)
	})

	// Teste de autorização quando o token é inválido
	t.Run("CustomAuthorizerHandler_InvalidToken", func(t *testing.T) {
		token := "invalid_token"
		methodArn := "arn:aws:execute-api:region:account-id:api-id/stage/method/resource-path"

		req := events.APIGatewayCustomAuthorizerRequest{
			AuthorizationToken: token,
			MethodArn:          methodArn,
		}

		response, err := controladores.CustomAuthorizerHandler(context.Background(), req, consultarClienteUC, autorizarUsuarioUC)
		assert.NoError(t, err)
		assert.Equal(t, "anonymous", response.PrincipalID)
		assert.Equal(t, "Deny", response.PolicyDocument.Statement[0].Effect)
	})
}
