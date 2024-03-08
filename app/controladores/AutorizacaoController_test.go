package controladores_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

// MockConsultarCliente simula o caso de uso ConsultarCliente.
type MockConsultarClienteAutorizacao struct{}

// ConsultarCliente simula a operação de consultar um cliente pelo ID.
func (m *MockConsultarCliente) MockConsultarClienteAutorizacao(idCliente string) (*dominio.Cliente, error) {
	// Simula a consulta do cliente pelo ID, retornando um cliente mockado
	return &dominio.Cliente{
		ID:     "123",
		CPF:    "12345678900",
		Nome:   "John Doe",
		Email:  "john.doe@example.com",
		Status: "ATIVO",
	}, nil
}

// MockAutorizarUsuario simula o caso de uso AutorizarUsuario.
type MockAutorizarUsuario struct{}

// AutorizarCliente simula a operação de autorizar um cliente.
func (m *MockAutorizarUsuario) AutorizarCliente(tokenString string) (bool, string) {
	// Simula a autorização do cliente, retornando um valor de autorização e o ID do usuário
	return true, "123"
}

func TestAutorizarcaoController_Handle(t *testing.T) {
	// Criando instâncias dos mocks dos casos de uso
	mockConsultarCliente := &MockConsultarCliente{}
	mockAutorizarUsuario := &MockAutorizarUsuario{}

	// Criando instância do controlador com os mocks dos casos de uso
	controller := controladores.NewAutorizarcaoController(mockConsultarCliente, mockAutorizarUsuario)

	// Executando a função a ser testada
	response, err := controller.Handle("mocked_token", "mocked_method_arn")

	// Verificando se não ocorreu nenhum erro
	assert.NoError(t, err)

	// Verificando se a resposta da autorização foi gerada corretamente
	assert.Equal(t, "123", response.PrincipalID)
	assert.Equal(t, "2012-10-17", response.PolicyDocument.Version)
	assert.Len(t, response.PolicyDocument.Statement, 1)
	assert.Equal(t, "Allow", response.PolicyDocument.Statement[0].Effect)
	assert.Equal(t, "execute-api:Invoke", response.PolicyDocument.Statement[0].Action[0])
	assert.Equal(t, "mocked_method_arn", response.PolicyDocument.Statement[0].Resource[0])
}
