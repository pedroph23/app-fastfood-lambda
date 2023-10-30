package apresentacao

type ClienteDTO struct {
	CPF   string `json:"cpf"`
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func NewClienteDTO(id, cpf, nome, email string) *ClienteDTO {
	return &ClienteDTO{
		CPF:   cpf,
		ID:    id,
		Nome:  nome,
		Email: email,
	}
}
