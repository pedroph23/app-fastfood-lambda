// usecases/autenticacao_cliente_uc.go
package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pedroph23/app-fastfood-lambda/app/dominio"
)

type AutenticarUsuarioImpl struct{}

func NewAutenticarUsuarioImpl() *AutenticarUsuarioImpl {
	return &AutenticarUsuarioImpl{}
}

func (uc *AutenticarUsuarioImpl) AutenticarClienteAnonimo() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("none"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = "anonimo"
	claims["iss"] = "appfastfood"

	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

	if err != nil {
		return "", fmt.Errorf("failed to create unsigned token: %v", err)
	}

	return tokenString, nil
}

func (uc *AutenticarUsuarioImpl) AutenticarCliente(cliente *dominio.Cliente) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("none"))
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = cliente.ID
	claims["iss"] = "appfastfood"

	tokenString, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	if err != nil {
		return "", fmt.Errorf("failed to create unsigned token: %v", err)
	}

	return tokenString, nil
}
