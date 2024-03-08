package controladores_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

// MockConsultarCliente simula o caso de uso ConsultarCliente.
type MockConsultarClienteAutenticacao struct{}

// ConsultarCliente simula a operação de consultar um cliente pelo ID.
func (m *MockConsultarClienteAutenticacao) ConsultarCliente(idCliente string) (*dominio.Cliente, error) {
	// Simula a consulta do cliente pelo ID, retornando um cliente mockado
	return &dominio.Cliente{
		ID:     "123",
		CPF:    "12345678900",
		Nome:   "John Doe",
		Email:  "john.doe@example.com",
		Status: "ATIVO",
	}, nil
}

// MockAutenticarUsuario simula o caso de uso AutenticarUsuario.
type MockAutenticarUsuario struct{}

// AutenticarClienteAnonimo simula a operação de autenticar um cliente anônimo.
func (m *MockAutenticarUsuario) AutenticarClienteAnonimo() (string, error) {
	// Simula a autenticação de um cliente anônimo, retornando um token
	return "mocked_anonymous_token", nil
}

// AutenticarCliente simula a operação de autenticar um cliente.
func (m *MockAutenticarUsuario) AutenticarCliente(cliente *dominio.Cliente) (string, error) {
	// Simula a autenticação de um cliente, retornando um token
	return "mocked_authenticated_token", nil
}

func TestAutenticacaoController_Handle(t *testing.T) {
	// Criando instâncias dos mocks dos casos de uso
	mockConsultarCliente := &MockConsultarCliente{}
	mockAutenticarUsuario := &MockAutenticarUsuario{}

	// Criando instância do controlador com os mocks dos casos de uso
	controller := controladores.NewAutenticacaoController(mockConsultarCliente, mockAutenticarUsuario)

	// Executando a função a ser testada para autenticar um cliente anônimo
	responseAnonymous, err := controller.Handle("anonimo")

	// Verificando se não ocorreu nenhum erro
	assert.NoError(t, err)
	// Verificando se o token de cliente anônimo foi gerado corretamente
	assert.Equal(t, "mocked_anonymous_token", string(responseAnonymous))

	// Executando a função a ser testada para autenticar um cliente autenticado
	responseAuthenticated, err := controller.Handle("123")

	// Verificando se não ocorreu nenhum erro
	assert.NoError(t, err)
	// Verificando se o token de cliente autenticado foi gerado corretamente
	assert.Equal(t, "mocked_authenticated_token", string(responseAuthenticated))
}
