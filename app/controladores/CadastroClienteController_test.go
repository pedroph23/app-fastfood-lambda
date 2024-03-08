package controladores_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/apresentacao"
	"github.com/pedroph23/app-fastfood-lambda/app/controladores"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

// MockCadastrarCliente simula o caso de uso CadastrarCliente.
type MockCadastrarCliente struct{}

// CadastrarCliente simula a operação de cadastrar um cliente.
func (m *MockCadastrarCliente) CadastrarCliente(input apresentacao.ClienteDTO) (*dominio.Cliente, error) {
	// Simula o cadastro do cliente, retornando um cliente mockado
	return &dominio.Cliente{
		ID:     "123",
		CPF:    input.CPF,
		Nome:   input.Nome,
		Email:  input.Email,
		Status: "ATIVO",
	}, nil
}

func TestCadastroClienteController_Handle(t *testing.T) {
	// Criando uma instância do mock do caso de uso
	mockCadastrarCliente := &MockCadastrarCliente{}

	// Criando instância do controlador com o mock do caso de uso
	controller := controladores.NewCadastroClienteController(mockCadastrarCliente)

	// Simulando um corpo de requisição contendo os dados do cliente
	requestBody := `{"cpf": "12345678900", "nome": "John Doe", "email": "john.doe@example.com"}`

	// Executando a função a ser testada
	response, err := controller.Handle(requestBody)

	// Verificando se não ocorreu nenhum erro
	assert.NoError(t, err)

	// Verificando se a resposta do cadastro foi gerada corretamente
	expectedResponse := map[string]string{
		"message":    "Cliente cadastrado com sucesso",
		"id_cliente": "123",
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)
	assert.Equal(t, expectedResponseBytes, response)
}
