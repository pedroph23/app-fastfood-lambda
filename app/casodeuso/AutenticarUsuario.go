// usecases/autenticacao_cliente_uc.go
package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type AutenticarUsuario struct{}

func NewAutenticarUsuario() *AutenticarUsuario {
	return &AutenticarUsuario{}
}

func (uc *AutenticarUsuario) AutenticarCliente(cliente *dominio.Cliente) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("none"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = cliente.ID
	claims["cpf"] = cliente.CPF
	claims["nome"] = cliente.Nome
	claims["email"] = cliente.Email
	claims["iss"] = "appfastfood"

	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		return "", fmt.Errorf("failed to create unsigned token: %v", err)
	}

	return tokenString, nil
}
