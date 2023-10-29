// usecases/autenticacao_cliente_uc.go
package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type AutenticarUsuario struct{}

func AutenticarUsuario() *AutenticarUsuario {
	return &AutenticarUsuario{}
}

func (uc *AutenticarUsuario) AutenticarCliente(cliente *domain.Cliente) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
		"user":  cliente.ID,
		"cpf":   cliente.CPF,
		"nome":  cliente.Nome,
		"email": cliente.Email,
		"iss":   "appfastfood",
	})

	tokenString, err := token.SignedString([]byte(""))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}
