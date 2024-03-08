package casodeuso_test

import (
	"testing"

	"github.com/pedroph23/app-fastfood-lambda/app/casodeuso"
	"github.com/stretchr/testify/assert"
)

func TestAutorizarUsuario_AutorizarCliente(t *testing.T) {
	uc := casodeuso.NewAutorizarUsuarioImpl()
	t.Run("TokenValido", func(t *testing.T) {
		tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlciI6IjEyMzQ1Njc4OTAiLCJpc3MiOiJhcHBmYXN0Zm9vZCIsImV4cCI6MTYxNzE1Mjg5MH0.CTf62E8zXQa7nvnyU4NRTtQcAB7fb0lhL5WpXWqATnU"

		ok, userID := uc.AutorizarCliente(tokenString)

		assert.True(t, ok)
		assert.Equal(t, "1234567890", userID)
	})

	t.Run("TokenInvalido", func(t *testing.T) {
		tokenString := "token_invalido"

		ok, userID := uc.AutorizarCliente(tokenString)
		assert.False(t, ok)
		assert.Empty(t, userID)
	})

	t.Run("TokenComEmissaoIncorreta", func(t *testing.T) {
		tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiaXNzIjoiZWFtb25zIiwiaWF0IjoxNjE0MDA0OTgwfQ.6as_6v80ziBdTp4I_hEwyrFpGyODlU-3Ly4YQoTRdOg"

		ok, userID := uc.AutorizarCliente(tokenString)
		assert.False(t, ok)
		assert.Empty(t, userID)
	})
}
