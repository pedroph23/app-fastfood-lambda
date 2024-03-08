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
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

func TestHandler(t *testing.T) {
	// Mock do repositório do cliente para simular o banco de dados
	clienteRepository := repositorio.NewRepositorioClienteMock()

	// Instância dos casos de uso
	autenticacaoClienteUC := casodeuso.NewAutenticarUsuarioImpl()
	consultarClienteUC := casodeuso.NewConsultarClienteImpl(clienteRepository)
	cadastrarClienteUC := casodeuso.NewCadastrarClienteImpl(clienteRepository)
	atualizarClienteUC := casodeuso.NewAtualizarClienteImpl(clienteRepository)
	autorizarUsuarioUC := casodeuso.NewAutorizarUsuarioImpl()

	t.Run("AutenticacaoClienteHandler", func(t *testing.T) {
		requestBody := `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "ATIVO"}`

		req := events.APIGatewayProxyRequest{
			HTTPMethod:     "POST",
			Path:           "/auth",
			PathParameters: map[string]string{"id_cliente": "123"},
			Body:           requestBody,
		}

		response, err := AutenticacaoClienteHandler(context.Background(), req, autenticacaoClienteUC, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		var authDTO apresentacao.AuthDTO
		err = json.Unmarshal([]byte(response.Body), &authDTO)
		assert.NoError(t, err)

	})

	t.Run("CadastroClienteHandler", func(t *testing.T) {
		requestBody := `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "ATIVO"}`

		req := events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/clientes",
			Body:       requestBody,
		}

		response, err := CadastroClienteHandler(context.Background(), req, cadastrarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

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

		response, err := ConsultaClienteHandler(context.Background(), req, consultarClienteUC)
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

		req := events.APIGatewayProxyRequest{
			HTTPMethod: "PATCH",
			Path:       "/clientes/123",
			PathParameters: map[string]string{
				"id_cliente": "123",
			},
			Body: requestBody,
		}

		response, err := AtualizaClienteHandler(context.Background(), req, atualizarClienteUC, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

	})

	// Teste de autorização
	t.Run("CustomAuthorizerHandler", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4SfMmK2KevnSRjhZB0G54n8iTW4hnRzEePzMOMxeMBw"
		methodArn := "arn:aws:execute-api:region:account-id:api-id/stage/method/resource-path"

		req := events.APIGatewayCustomAuthorizerRequest{
			AuthorizationToken: token,
			MethodArn:          methodArn,
		}

		response, err := CustomAuthorizerHandler(context.Background(), req, consultarClienteUC, autorizarUsuarioUC)
		assert.NoError(t, err)
		assert.NotNil(t, response.PrincipalID)
	})

	// Teste de autorização quando o token é inválido
	t.Run("CustomAuthorizerHandler_InvalidToken", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4SfMmK2KevnSRjhZB0G54n8iTW4hnRzEePzMOMxeMBw"
		methodArn := "arn:aws:execute-api:region:account-id:api-id/stage/method/resource-path"

		req := events.APIGatewayCustomAuthorizerRequest{
			AuthorizationToken: token,
			MethodArn:          methodArn,
		}

		response, err := CustomAuthorizerHandler(context.Background(), req, consultarClienteUC, autorizarUsuarioUC)
		assert.NoError(t, err)
		assert.Equal(t, "anonymous", response.PrincipalID)
		assert.Equal(t, "Deny", response.PolicyDocument.Statement[0].Effect)
	})

	t.Run("AutenticacaoClienteHandler_Sucesso", func(t *testing.T) {
		// Envie um token válido para autenticar um cliente
		req := events.APIGatewayProxyRequest{
			HTTPMethod:     "POST",
			Path:           "/auth",
			PathParameters: map[string]string{"id_cliente": "123"},
		}

		response, err := AutenticacaoClienteHandler(context.Background(), req, autenticacaoClienteUC, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("CadastroClienteHandler_Sucesso", func(t *testing.T) {
		// Envie um corpo de requisição válido para cadastrar um cliente
		req := events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Path:       "/clientes",
			Body:       `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "ATIVO"}`,
		}

		response, err := CadastroClienteHandler(context.Background(), req, cadastrarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("ConsultaClienteHandler_Sucesso", func(t *testing.T) {
		// Consulte um cliente existente na base de dados
		req := events.APIGatewayProxyRequest{
			HTTPMethod:     "GET",
			Path:           "/clientes/123",
			PathParameters: map[string]string{"id_cliente": "123"},
		}

		response, err := ConsultaClienteHandler(context.Background(), req, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("AtualizaClienteHandler_Sucesso", func(t *testing.T) {
		// Envie um corpo de requisição válido para atualizar um cliente existente na base de dados
		req := events.APIGatewayProxyRequest{
			HTTPMethod: "PATCH",
			Path:       "/clientes/123",
			PathParameters: map[string]string{
				"id_cliente": "123",
			},
			Body: `{"cpf": "12345678900", "id": "123", "nome": "John Doe", "email": "john.doe@example.com", "status": "INATIVO"}`,
		}

		response, err := AtualizaClienteHandler(context.Background(), req, atualizarClienteUC, consultarClienteUC)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})
}
