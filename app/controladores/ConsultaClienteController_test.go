package controladores_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

// MockConsultarCliente simula o caso de uso ConsultarCliente.
type MockConsultarClienteController struct{}

// ConsultarCliente simula a operação de consultar um cliente pelo ID.
func (m *MockConsultarClienteController) ConsultarCliente(idCliente string) (*dominio.Cliente, error) {
	// Simula a busca do cliente pelo ID, retornando um cliente mockado
	return &dominio.Cliente{
		ID:     "123",
		CPF:    "12345678900",
		Nome:   "John Doe",
		Email:  "john.doe@example.com",
		Status: "ATIVO",
	}, nil
}

func TestConsultaClienteController_Handle(t *testing.T) {
	// Criando uma instância do mock do caso de uso
	mockConsultarCliente := &MockConsultarCliente{}

	// Criando instância do controlador com o mock do caso de uso
	controller := controladores.NewConsultaClienteController(mockConsultarCliente)

	// Simulando o ID do cliente a ser consultado
	idCliente := "123"

	// Executando a função a ser testada
	cliente, err := controller.Handle(idCliente)

	// Verificando se não ocorreu nenhum erro
	assert.NoError(t, err)

	// Verificando se o cliente retornado corresponde ao esperado
	expectedCliente := &dominio.Cliente{
		ID:     "123",
		CPF:    "12345678900",
		Nome:   "John Doe",
		Email:  "john.doe@example.com",
		Status: "INATIVO",
	}
	assert.Equal(t, expectedCliente, cliente)
}
