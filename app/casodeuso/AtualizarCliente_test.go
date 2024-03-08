package casodeuso_test

import (
	"testing"

	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
	"github.com/pedroph23/app-fastfood-lambda/app/repositorio"
	"github.com/stretchr/testify/assert"
)

// Mock do repositório do cliente para os testes

func TestAtualizarCliente_AtualizarCliente(t *testing.T) {
	repositorioMock := repositorio.NewRepositorioClienteMock()
	uc := casodeuso.NewAtualizarClienteImpl(repositorioMock)

	t.Run("AtualizarCliente_Ativo", func(t *testing.T) {
		inputCliente := &dominio.Cliente{
			CPF:    "12345678900",
			ID:     "123",
			Nome:   "John Doe",
			Email:  "john.doe@example.com",
			Status: "ATIVO",
		}

		expectedID := "39b5177e82858ecc5661a2077b58edc3"
		expectedStatus := "ATIVO"

		clienteAtualizado, err := uc.AtualizarCliente(inputCliente)
		assert.NoError(t, err)
		assert.NotNil(t, clienteAtualizado)
		assert.Equal(t, expectedID, clienteAtualizado.ID)
		assert.Equal(t, expectedStatus, clienteAtualizado.Status)
	})

	t.Run("AtualizarCliente_Inativo", func(t *testing.T) {
		inputCliente := &dominio.Cliente{
			CPF:    "12345678900",
			ID:     "123",
			Nome:   "John Doe",
			Email:  "john.doe@example.com",
			Status: "INATIVO",
		}

		expectedID := "123" // MD5 hash of "12345678900"
		expectedStatus := "INATIVO"

		clienteAtualizado, err := uc.AtualizarCliente(inputCliente)
		assert.NoError(t, err)
		assert.NotNil(t, clienteAtualizado)
		assert.Equal(t, expectedID, clienteAtualizado.ID)
		assert.Equal(t, expectedStatus, clienteAtualizado.Status)
	})

}
