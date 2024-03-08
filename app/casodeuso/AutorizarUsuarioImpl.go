package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type AutorizarUsuarioImpl struct {
}

func NewAutorizarUsuarioImpl() *AutorizarUsuarioImpl {
	return &AutorizarUsuarioImpl{}
}

func (uc *AutorizarUsuarioImpl) AutorizarCliente(tokenString string) (bool, string) {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		fmt.Println("Erro ao fazer parse do token: ", err)
		return false, ""
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	if claims["iss"] != "appfastfood" {
		// Se os requisitos não forem atendidos, bloqueie a requisição]
		fmt.Println("Bloqueando requisição")
		return false, ""

	}

	return true, claims["user"].(string)
}
