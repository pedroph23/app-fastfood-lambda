package casodeuso

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type AutorizarUsuario struct {
}

func NewAutorizarUsuario() *AutorizarUsuario {
	return &AutorizarUsuario{}
}

func (uc *AutorizarUsuario) AutorizarCliente(tokenString string) (bool, string) {

	if tokenString == "" {
		return true, ""
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})

	if err != nil {
		fmt.Println("Erro ao fazer parse do token: ", err)
		return false, ""
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	if claims["iss"] != "appfastfood" || claims["user"] == "" {
		// Se os requisitos não forem atendidos, bloqueie a requisição]
		fmt.Println("Bloqueando requisição")
		return false, ""
	}

	return true, claims["user"].(string)
}
