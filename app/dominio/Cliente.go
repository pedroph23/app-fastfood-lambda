package dominio

type Cliente struct {
	CPF   string
	ID    string
	Nome  string
	Email string
}

func NewCliente(cpf, id, nome, email string) *Cliente {
	return &Cliente{
		CPF:   cpf,
		ID:    id,
		Nome:  nome,
		Email: email,
	}
}
