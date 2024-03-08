package casodeuso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

func TestConsultarCliente_ConsultarCliente(t *testing.T) {
	// Crie o mock do repositório do cliente
	clienteRepository := repositorio.NewRepositorioClienteMock()

	// Crie uma instância do caso de uso
	uc := casodeuso.NewConsultarClienteImpl(clienteRepository)

	t.Run("ClienteExistente", func(t *testing.T) {
		// ID do cliente existente no mock do repositório
		idCliente := "123"

		cliente, err := uc.ConsultarCliente(idCliente)
		assert.NoError(t, err)
		assert.NotNil(t, cliente)
		assert.Equal(t, "12345678900", cliente.CPF)
		assert.Equal(t, "John Doe", cliente.Nome)
		assert.Equal(t, "john.doe@example.com", cliente.Email)
		assert.Equal(t, "ATIVO", cliente.Status)
	})

}
