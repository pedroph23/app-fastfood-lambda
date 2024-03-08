package controladores_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

// MockAtualizarCliente simula o caso de uso AtualizarClienteImpl.
type MockAtualizarCliente struct{}

// AtualizarCliente simula a operação de atualizar um cliente.
func (m *MockAtualizarCliente) AtualizarCliente(inputCliente *dominio.Cliente) (*dominio.Cliente, error) {
	// Simula a atualização do cliente, retornando o cliente atualizado
	inputCliente.Status = "ATIVO"
	return inputCliente, nil
}

// MockConsultarCliente simula o caso de uso ConsultarClienteImpl.
type MockConsultarCliente struct{}

// ConsultarCliente simula a operação de consultar um cliente pelo ID.
func (m *MockConsultarCliente) ConsultarCliente(idCliente string) (*dominio.Cliente, error) {
	// Simula a consulta do cliente pelo ID, retornando um cliente mockado
	return &dominio.Cliente{
		ID:     "123",
		CPF:    "12345678900",
		Nome:   "John Doe",
		Email:  "john.doe@example.com",
		Status: "INATIVO",
	}, nil
}

func TestAtualizarClienteController_Handle(t *testing.T) {
	// Criando instâncias dos mocks dos casos de uso
	mockAtualizarCliente := &MockAtualizarCliente{}
	mockConsultarCliente := &MockConsultarCliente{}

	// Criando instância do controlador com os mocks dos casos de uso
	controller := controladores.NewAtualizarClienteController(mockAtualizarCliente, mockConsultarCliente)

	// Dados de entrada para o teste
	idCliente := "123"
	requestBody := `{"id": "123", "cpf": "12345678900", "nome": "John Doe", "email": "john.doe@example.com", "status": "INATIVO"}`

	// Executando a função a ser testada
	response, err := controller.Handle(idCliente, requestBody)

	// Verificando se não ocorreu nenhum erro
	assert.NoError(t, err)

	// Verificando se o cliente foi atualizado com sucesso
	var responseBody map[string]string
	err = json.Unmarshal(response, &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Cliente atualizado com sucesso", responseBody["message"])
	assert.Equal(t, "123", responseBody["id_cliente"])
}
