package apresentacao

type ClienteDTO struct {
	CPF   string
	ID    string
	Nome  string
	Email string
}

func NewClienteDTO(cpf, id, nome, email string) *ClienteDTO {
	return &ClienteDTO{
		CPF:   cpf,
		ID:    id,
		Nome:  nome,
		Email: email,
	}
}
