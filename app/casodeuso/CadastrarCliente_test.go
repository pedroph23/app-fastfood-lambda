package casodeuso_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedroph23/app-fastfood-lambda/app/apresentacao"
	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
)

func TestCadastrarCliente_CadastrarCliente(t *testing.T) {
	// Crie o mock do repositório do cliente
	clienteRepository := repositorio.NewRepositorioClienteMock()

	// Crie uma instância do caso de uso
	uc := casodeuso.NewCadastrarClienteImpl(clienteRepository)

	t.Run("ClienteNovo", func(t *testing.T) {
		inputCliente := apresentacao.ClienteDTO{
			CPF:   "12345678900",
			Nome:  "John Doe",
			Email: "john.doe@example.com",
		}

		cliente, err := uc.CadastrarCliente(inputCliente)
		assert.NoError(t, err)
		assert.NotNil(t, cliente)
		assert.Equal(t, "12345678900", cliente.CPF)
		assert.Equal(t, "John Doe", cliente.Nome)
		assert.Equal(t, "john.doe@example.com", cliente.Email)
		assert.Equal(t, "ATIVO", cliente.Status)
	})

	t.Run("ClienteExistente", func(t *testing.T) {
		inputCliente := apresentacao.ClienteDTO{
			CPF:   "12345678900",
			Nome:  "Jane Doe",
			Email: "jane.doe@example.com",
		}

		cliente, err := uc.CadastrarCliente(inputCliente)
		assert.NoError(t, err)
		assert.NotNil(t, cliente)
		assert.Equal(t, "12345678900", cliente.CPF)
		assert.Equal(t, "Jane Doe", cliente.Nome)
		assert.Equal(t, "jane.doe@example.com", cliente.Email)
		assert.Equal(t, "ATIVO", cliente.Status)
	})
}
