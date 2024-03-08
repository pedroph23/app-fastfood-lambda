package casodeuso

// AutorizarUsuario é a interface que define o caso de uso de autorização de usuário
type AutorizarUsuario interface {
	AutorizarCliente(tokenString string) (bool, string)
}
